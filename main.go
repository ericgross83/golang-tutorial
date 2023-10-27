package main

import "fmt"

func main() {
	// x := 0
	// for x < 5 {
	// 	fmt.Println("value of x is:", x+1)
	// 	x++
	// }

	// for i := 0; i < 5; i++ {
	// 	fmt.Println("value of i is:", i+1)
	// }

	names := []string{"yoshi", "mario", "luigi", "peach"}

	// for i := 0; i < len(names); i++ {
	// 	fmt.Println(names[i])
	// }

	// for index, name := range names {
	// 	fmt.Printf("The value at index %v is %v\n", index, name)
	// }

	for _, name := range names {
		fmt.Printf("The value is %v \n", name)
	}

}
