package main

// Разработать программу, которая проверяет, что все символы в строке уникальные (true — если уникальные, false etc). Функция проверки должна быть регистронезависимой.

import (
	"fmt"
	"strings"
)

func IsUnique(s string) bool {
	lowercaseStr := strings.ToLower(s)

	charMap := make(map[rune]bool)

	for _, char := range lowercaseStr {
		if charMap[char] {
			return false
		}
		charMap[char] = true
	}

	return true
}

func main() {
	str1 := "abcd"
	str2 := "abCdefAaf"
	str3 := "aabcd"

	fmt.Println(IsUnique(str1)) // true
	fmt.Println(IsUnique(str2)) // false
	fmt.Println(IsUnique(str3)) // false
}