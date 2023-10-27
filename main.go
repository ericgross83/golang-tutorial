package main

import "fmt"

func main() {

	menu := map[string]float64{
		"soup":           4.99,
		"pie":            7.99,
		"salad":          6.99,
		"toffee pudding": 3.55,
	}

	fmt.Println(menu)
	fmt.Println(menu["pie"])

	// looping maps
	for key, item := range menu {
		fmt.Println(key, "-", item)
	}

	// ints as key type
	phonebook := map[int]string{
		267584967: "mario",
		267584154: "luigi",
		267584758: "peach",
	}

	fmt.Println(phonebook)
	fmt.Println(phonebook[267584967])

	phonebook[267584967] = "bowser"

	fmt.Println(phonebook[267584967])
}
