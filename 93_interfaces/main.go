package main

import "fmt"

type Speaker interface {
	SayHello()
}

type Human struct {
	Name string
}

func (h Human) SayHello() {
	fmt.Println("Hello from", h.Name)
}

func (h Human) Greet() {
	fmt.Println(h.Name + "is greeting you")
}

func DoGreeting(actor Speaker) {
	actor.SayHello()
}

func main() {
	someHuman := Human{"Name Vorname"}

	DoGreeting(someHuman)
}
