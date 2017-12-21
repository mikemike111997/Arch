package main

import (
	"fmt"
	"./set"
	"strconv"
	// "reflect"
	"time"
	"runtime"

)

func main() {

	// Initialize Set
	set1 := set.New()

	// Add items concurrently (item1, item2, and so on)
	for i := 0; i < 1000; i++ {
		item := "item" + strconv.Itoa(i)
		set1.Add(item)
	}


	// fmt.Println("First set (set1):\n\t", set1)
	// fmt.Println("TypeOf(set1) = ", reflect.TypeOf(set1).String())

	set2 := set.New()
	for i := 500; i < 1500; i++ {
		item := "item" + strconv.Itoa(i)
		set2.Add(item)
	}
	// fmt.Println("Second set (set2):\n\t", set2)
	// fmt.Println("TypeOf(set2) = ", reflect.TypeOf(set2).String())

	t0 := time.Now();
	// fmt.Println("set1 + set2:\n\t", set.Union(set1, set2))
	// fmt.Println("set1 - set2\n\t", set.Difference(set1, set2))
	// fmt.Println("set1 * set2\n\t", set.Intersection(set1, set2))
	
	var _ = runtime.GOMAXPROCS(30)
	fmt.Println(runtime.GOMAXPROCS(-1))

	set.Union(set1, set2)
	set.Difference(set1, set2)
	set.Intersection(set1, set2)

	t1 := time.Now();
    fmt.Printf("Elapsed time: %v", t1.Sub(t0));





}