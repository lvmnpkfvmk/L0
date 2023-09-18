package main

// Реализовать собственную функцию sleep.

import (
	"fmt"
	"time"
)

func Sleep(milliseconds int) {
	select {
	case <-time.After(time.Duration(milliseconds) * time.Millisecond):
		return
	}
}

func main() {
	fmt.Println("Start")
	Sleep(2000) // Lasts 2 seconds
	fmt.Println("End")
}