package main

// Реализовать быструю сортировку массива (quicksort) встроенными методами языка.

import (
	"fmt"
)

func partition(arr []int, low, high int) int {
	pivot := arr[high]

	i := low - 1

	for j := low; j < high; j++ {
		if arr[j] <= pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}

	arr[i+1], arr[high] = arr[high], arr[i+1]
	return i + 1
}

func quicksort(arr []int, low, high int) {
	if low < high {
		pivotIndex := partition(arr, low, high)

		quicksort(arr, low, pivotIndex-1)
		quicksort(arr, pivotIndex+1, high)
	}
}

func main() {
	arr := []int{1, 6, 5, 1111, 32, 173, 3, 10, 3, 175, 346} 
	quicksort(arr, 0, len(arr)-1)

	fmt.Println("Sorted:", arr) // [1 3 3 5 6 10 32 173 175 346 1111]
}