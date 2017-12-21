package set

import (
	"sync"
)

// Interface is describing a Set. Sets are an unordered, unique list of values.
type Interface interface {
	New(items ...interface{}) Interface
	Add(items ...interface{})
	Remove(items ...interface{})
	Pop() interface{}
	Has(items ...interface{}) bool
	Size() int
	Clear()
	IsEmpty() bool
	IsEqual(s Interface) bool
	IsSubset(s Interface) bool
	IsSuperset(s Interface) bool
	Each(func(interface{}) bool)
	String() string
	List() []interface{}
	Copy() Interface
	Merge(s Interface)
	Separate(s Interface)
}

// helpful to not write everywhere struct{}{}
var keyExists = struct{}{}

// Provides a common set.
type set struct {
	m map[interface{}]struct{} // struct{} doesn't take up space
}

// Union is the merger of multiple sets. It returns a new set with all the
// elements present in all the sets that are passed.
//
// The dynamic type of the returned set is determined by the first passed set's
// implementation of the New() method.
func Union(sets ...Interface) Interface {
	u := New()
	for _, set := range sets {
		set.Each(func(item interface{}) bool {
			u.Add(item)
			return true
		})
	}

	return u
}

// Difference returns a new set which contains items which are in in the first
// set but not in the others. Unlike the Difference() method you can use this
// function separately with multiple sets.
func Difference(set1, set2 Interface, sets ...Interface) Interface {
	s := set1.Copy()
	s.Separate(set2)
	for _, set := range sets {
		s.Separate(set) // seperate is thread safe
	}
	return s
}

// Intersection returns a new set which contains items that only exist in all given sets.
func Intersection(sets ...Interface) Interface {
	all := Union(sets...)
	result := Union(sets...)
	var wg sync.WaitGroup
	all.Each(func(item interface{}) bool {
	// 	if !set1.Has(item) || !set2.Has(item) {
	// 		result.Remove(item)
	// 	}
		

		for _, set := range sets {
			// fmt.Println(item)
			wg.Add(1)
			go func(s interface{}, wg *sync.WaitGroup){

				defer wg.Done()

				if !set.Has(item) {
					result.Remove(item)
				}
				
			}(set, &wg)

		}
		
		return true
	})

	wg.Wait()
	return result
}