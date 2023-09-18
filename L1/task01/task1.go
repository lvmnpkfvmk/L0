package main

//Дана структура Human (с произвольным набором полей и методов). Реализовать встраивание методов в структуре Action от родительской структуры Human (аналог наследования).

import "fmt"

type Human struct {
	age  int
	name string
}

func (h Human) Greeting() string {
	return fmt.Sprintf("Hello, my name is %s", h.name)
}

func (h Human) Get() string {
	return fmt.Sprintf("My age is %d", h.age)
}

type Action struct {
	Human
	action string
}

func (a *Action) Get() string {
	return fmt.Sprintf("The action is %s", a.action)
}

func main() {
	a := Action{Human{12, "John"}, "Eat"}
	fmt.Println(a.Greeting())  // a.Human.greeting()
	fmt.Println(a.Get())       // The action is Eat
	fmt.Println(a.Human.Get()) // My age is 12
}
