package main

import "fmt"

type example struct {
	flag    bool
	counter int16
	pi      float32
}

func main() {

	//Declare a var of type example set to its default values
	var e1 example

	fmt.Printf("%+v\n", e1)

	//Declare and init
	e2 := example{
		flag:    true,
		counter: 10,
		pi:      3.14,
	}

	//Display fields
	fmt.Println("Flag", e2.flag)
	fmt.Println("Counter", e2.counter)
	fmt.Println("Pi", e2.pi)

	//short init
	e3 := example{false, 10, 13.5}
	fmt.Println(e3)

}
