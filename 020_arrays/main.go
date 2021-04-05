package main

import "fmt"

const size uint = 3

func main() {
	//init array with all default values
	var a1 [3]int
	fmt.Println("array", a1, "length", len(a1))

	//size as const
	var a2 [2 * size]bool
	fmt.Println(a2, "len", len(a2))

	//size by it values
	a3 := [...]int{1, 2, 3}
	fmt.Println("len during init", a3, "len", len(a3))

	//change value
	a3[1] = 12
	fmt.Println("a3 after change", a3)

	//array out of bounds at compile time
	//a3[4] = 12
	// invalid array index 4 (out of bounds for 3-element array)

	//matrix
	var aa [3][3]int
	aa[1][1] = 1
	fmt.Println("Matrix", aa)

}
