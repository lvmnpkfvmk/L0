package main

//Имеется последовательность строк - (cat, cat, dog, cat, tree) создать для нее собственное множество.

import (
	"fmt"
)

func main() {
	sequence := []string{"cat", "cat", "dog", "cat", "tree"}
	set := make(map[string]bool)

	for _, str := range sequence {
		set[str] = true
	}

	for str := range set {
		fmt.Println(str) // cat, dog, tree (order unspecified)
	}
}
