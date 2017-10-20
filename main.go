package main

import (
	"fmt"
	"./set"
	"strconv"
	"sync"
)

func main() {
	var wg sync.WaitGroup // this is just for waiting until all goroutines finish

	// Initialize our thread safe Set
	s := set.New()

	// Add items concurrently (item1, item2, and so on)
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			item := "item" + strconv.Itoa(i)
			fmt.Println("adding", item)
			s.Add(item)
			wg.Done()
		}(i)
	}

	// Wait until all concurrent calls finished and print our set
	wg.Wait()
	fmt.Println(s)
	s2 := set.New("new1", "new2", "new3")
	fmt.Println(s2)

	tmp := set.Union(s, s2)
	fmt.Println(tmp)
}