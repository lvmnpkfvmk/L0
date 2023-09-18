package main

// Разработать программу, которая в рантайме способна определить тип переменной: int, string, bool, channel из переменной типа interface{}.


import (
	"fmt"
)

func SayType(x interface{}) {
	switch x.(type) {
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	case bool:
		fmt.Println("bool")
	case chan int:
		fmt.Println("chan int")
	default:
		fmt.Println("unknown")
	}
}

func main() {
	var x interface{}
	x = 0
	SayType(x) // int

	x = "x"
	SayType(x) // string

	x = true
	SayType(x) // bool

	x = make(chan int)
	SayType(x) // chan int

}