package main

// Реализовать бинарный поиск встроенными методами языка.

import "fmt"

func binarySearch(arr []int, target int) (error, int) {
	low := 0
	high := len(arr) - 1

	for low <= high {
		mid := (low + high) / 2

		if arr[mid] == target {
			return nil, mid
		} else if arr[mid] < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return fmt.Errorf("Target not found"), -1
}

func main() {
	arr := []int{1, 3, 5, 6, 10, 32, 173, 175, 346, 1111}

	err, index := binarySearch(arr, 10)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(index) // 4
	}

	err, index = binarySearch(arr, 32)

	if err != nil {
		fmt.Println(index)
	} else {
		fmt.Println(err) // nil
	}
}