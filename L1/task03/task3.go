package main

// Дана последовательность чисел: 2,4,6,8,10. Найти сумму их квадратов(22+32+42….) с использованием конкурентных вычислений.

import (
	"fmt"
	"sync"
)

func main() {
	array := []int{2, 4, 6, 8, 10}
	var wg sync.WaitGroup
	ch := make(chan int, 5)

	for _, x := range array {
		wg.Add(1)
		go func(x int) {
			defer wg.Done()
			ch <- x * x
		}(x)
	}

	wg.Wait()
	close(ch)
	sum := 0
	for x := range ch {
		sum += x
	}
	fmt.Println(sum)
}
