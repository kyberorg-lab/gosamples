package main

import "fmt"

func main() {
	i := 5
	b := &i
	//var b *int
	*b = 6
	println(i)

	s1 := []int{0, 0, 0}
	fmt.Println(s1)
	Sum(s1)
	fmt.Println(s1)
}

func Sum(items []int) {
	items[2] = 1
}
