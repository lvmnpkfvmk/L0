package main

// Поменять местами два числа без создания временной переменной.

import "fmt"

func main() {
    a, b := 1, 2

	fmt.Println(a, b) // 1 2

	a, b = b, a

	fmt.Println(a, b) // 2 1
}