package main

import (
	"fmt"
	"sort"
)

func main() {
	// greeting := "hello there firends!"

	// fmt.Println(strings.Contains(greeting, "hello"))
	// fmt.Println(strings.ReplaceAll(greeting, "hello", "hi"))
	// fmt.Println(strings.ToUpper(greeting))
	// fmt.Println(strings.Index(greeting, "ll"))
	// fmt.Println(strings.Split(greeting, " "))

	// the original value is unchanged
	// fmt.Println("original string was: ", greeting)

	//Sort

	// ages := []int{45, 20, 35, 30, 75, 60, 50, 25}
	// sort.Ints(ages)
	// fmt.Println(ages)

	// index := sort.SearchInts(ages, 30)
	// fmt.Println(index)

	names := []string{"yoshi", "mario", "peach", "bowser", "luigi"}

	sort.Strings(names)
	fmt.Println(names)

}
