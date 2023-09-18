package main

import (
	"fmt"
	"strings"
)

var justString string

func createHugeString(strSize int) string {
	var sb strings.Builder
	for i := 0; i < strSize; i++ {
		sb.WriteString("x")
	}
	return sb.String()
}

func someFunc() {
	v := createHugeString(1 << 10)
	// вместо того, чтобы хранить ссылку на огромную строку, скопируем первые 100 байт, дабы избежать утечки памяти
	justString = string([]byte(v[:100]))
}


func main() {
	someFunc()
	fmt.Println(justString)
}