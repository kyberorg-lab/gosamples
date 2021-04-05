package main

import "fmt"

func main() {
	var i interface{} = "hello"

	//short method (when you are sure)
	s := i.(string)
	fmt.Println(s)

	//more polite check
	s1, ok := i.(string)
	fmt.Println(s1, ok)

	f, ok := i.(string)
	fmt.Println(f, ok)

	//f = i.(float64) //panic
	//fmt.Println(f)
}
