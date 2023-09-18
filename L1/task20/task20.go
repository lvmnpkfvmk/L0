package main

// Разработать программу, которая переворачивает слова в строке. Пример: «snow dog sun — sun dog snow».

import (
	"fmt"
	"strings"
)

func reverseWords(s string) string {
    words := strings.Fields(s)
    reverse(words)
    return strings.Join(words, " ")
}

func reverse(arr []string) {
    if len(arr) == 0 { 
		return
	}
    l, r := 0, len(arr) - 1
    for l < r {
        arr[l], arr[r] = arr[r], arr[l]
        l++
        r--
    }
}
func main() {
	str := "snow dog sun"
	fmt.Println(reverseWords(str)) // sun dog snow
}