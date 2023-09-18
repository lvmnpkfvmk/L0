package main

// Удалить i-ый элемент из слайса.

import (
	"fmt"
)

func removeElement(slice []int, index int) []int {
	if index < 0 || index >= len(slice) {
		return slice
	}
	return append(slice[:index], slice[index+1:]...)
}

func main() {
	slice := []int{1, 2, 3, 4, 5}
	fmt.Println("Removed:", removeElement(slice, 4)) // [1 2 3 4]
	fmt.Println("Removed:", removeElement(slice, 100)) // [1 2 3 4 5] nothing happens
}