package main

import (
	"fmt"
	"strings"
)

func getInitials(n string) (string, string) {
	upper := strings.ToUpper(n)
	names := strings.Split(upper, " ")

	var initials []string
	for _, value := range names {
		initials = append(initials, value[:1])
	}

	if len(initials) > 1 {
		return initials[0], initials[1]
	}

	return initials[0], "_"
}

func main() {
	initial1, initial2 := getInitials("tifa lockhart")
	fmt.Println(initial1, initial2)

	initial3, initial4 := getInitials("cloud strife")
	fmt.Println(initial3, initial4)

	initial5, initial6 := getInitials("barret")
	fmt.Println(initial5, initial6)
}
