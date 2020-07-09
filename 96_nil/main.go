package main

import "fmt"

func main() {
	var z interface{}

	fmt.Println("%v %v\n", z, z == nil)

	if f := getNil(10); f != nil {
		fmt.Println("I am not nill")
	}

	if f := getTrueNil(); f == nil {
		fmt.Println("But I actually am nil")
	}
}

type myError string

func (m myError) Error() string {
	return ""
}

func getNil(input interface{}) error {
	var m *myError
	if _, ok := input.(int); ok {
		return m
	}
	return nil
}

func getTrueNil() error {
	return nil
}
