package main

import "fmt"

func updateName(name string) {
	name = "wedge"
}

func main() {

	// group A types -> strings, ints, bools, floats, arrays, structs
	name := "tifa"

	updateName(name)

	fmt.Println(name)
}
