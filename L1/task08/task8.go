package main

// Дана переменная int64. Разработать программу которая устанавливает i-й бит в 1 или 0.

import (
	"fmt"
)

func setBit(num int64, i int, bit int) int64 {
	if bit == 1 {
		return (1 << i) | num
	} else {
		return (^(1 << i)) & num
	}

}

func main() {
	var num int64
	i := 2
	num = setBit(num, i, 1)
	fmt.Printf("%b\n", num) // 0b100
	num = setBit(num, i, 0)
	fmt.Printf("%b\n", num) // 0b0

}
