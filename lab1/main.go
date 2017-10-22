package lab1

import (
	"fmt"
	"./set"
	"strconv"
	"reflect"
)

func main() {

	// Initialize Set
	s := set.New()

	// Add items concurrently (item1, item2, and so on)
	for i := 0; i < 10; i++ {
		item := "item" + strconv.Itoa(i)
		fmt.Println("adding", item)
		s.Add(item)
	}

	fmt.Println(s)
	s2 := set.New("new1", "new2", "new3")
	fmt.Println(s2)

	tmp := set.Union(s, s2)
	fmt.Println(tmp)

	st := set.New("1", "2", "3")
	rt := set.New("3", "4", "5")
	xt := set.New("5", "6", "7")

	u := set.Union(st, rt, xt)
	settype := reflect.TypeOf(u).String()
	fmt.Println("result = ", u)
	fmt.Println(settype)
}