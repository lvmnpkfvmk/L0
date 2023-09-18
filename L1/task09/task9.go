package main

import (
	"fmt"
)

func main() {
	in := make(chan int)
	out := make(chan int)

	go func() {
		numbers := []int{1, 2, 3, 4, 5}
		for _, num := range numbers {
			in <- num
		}
		close(in)
	}()
	go func() {
		for x := range in {
			result := x * 2
			out <- result
		}
		close(out)
	}()
	for result := range out {
		fmt.Println(result)
	}
}
