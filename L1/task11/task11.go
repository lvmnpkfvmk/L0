package main

import "fmt"

//Реализовать пересечение двух неупорядоченных множеств.

func Intersect(a, b []int) []int {
	mp := make(map[int]bool, len(a))
	for _, x := range a {
		mp[x] = true
	}
	var ans []int
	for _, x := range b {
		if mp[x] {
			ans = append(ans, x)
		}
	}
	return ans
}

func main() {
	a := []int{1, 2, 3, 7}
	b := []int{2, 4, 3, 17}

	aintb := Intersect(a, b)

	fmt.Println("Intersected:")
	for _, item := range aintb {
		fmt.Println(item) // 2, 3
	}
}
