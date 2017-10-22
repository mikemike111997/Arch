package main

import (
	"fmt"
	"./set"
	"strconv"
	"reflect"
)

func main() {

	// Initialize Set
	set1 := set.New()

	// Add items concurrently (item1, item2, and so on)
	for i := 0; i < 5; i++ {
		item := "item" + strconv.Itoa(i)
		set1.Add(item)
	}

	fmt.Println("First set (set1):\n\t", set1)
	fmt.Println("TypeOf(set1) = ", reflect.TypeOf(set1).String())

	set2 := set.New("item0", "item3", "new1", "new2", "new3")
	fmt.Println("Second set (set2):\n\t", set2)
	fmt.Println("TypeOf(set2) = ", reflect.TypeOf(set2).String())

	fmt.Println("set1 + set2:\n\t", set.Union(set1, set2))
	fmt.Println("set1 - set2\n\t", set.Difference(set1, set2))
	fmt.Println("set1 * set2\n\t", set.Intersection(set1, set2))

}