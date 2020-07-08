package main

import "fmt"

func main() {
	var sl []int
	fmt.Println("Empty slice", sl)

	//adding element to slice
	sl = append(sl, 100)
	fmt.Println("Appended slice", sl)

	//len
	fmt.Println("Slice Len", len(sl))

	//capacity
	fmt.Println("capacity", sl, cap(sl))
	sl = append(sl, 101)
	fmt.Println("capacity", sl, cap(sl))
	sl = append(sl, 102)
	fmt.Println("capacity", sl, cap(sl))
	sl = append(sl, 103)
	sl = append(sl, 104)
	fmt.Println("capacity", sl, cap(sl))
	sl = append(sl, 105)
	fmt.Println("capacity", sl, cap(sl))

	//short init
	sl2 := []int{10, 20, 30}
	fmt.Println(sl2)

	//adding slice to slice
	sl = append(sl, sl2...)
	fmt.Println(sl)

	//slice matrix
	var slm [][]int
	slm = append(slm, sl2)
	fmt.Println(slm)

	//slice with init capacity
	slice3 := make([]int, 10)
	fmt.Println(slice3, len(slice3), cap(slice3))

	//slice with init size and capacity
	slice4 := make([]int, 0, 15)
	fmt.Println(slice4, len(slice4), cap(slice4))
	slice4 = append(slice4, []int{1, 2, 3, 4, 5, 6}...)
	fmt.Println(slice4, len(slice4), cap(slice4))

	//link to inner array is copied when do :=
	slice5 := slice4
	slice5[1] = 100500
	fmt.Println(slice4, slice5)

	slice4 = append(slice4, []int{1, 2, 3, 4, 5, 6}...)
	slice4[1] = 999
	fmt.Println(slice4, slice5)

	//wrong attempt to copy slice
	var slice6 []int
	copy(slice6, slice5)
	fmt.Println(slice6)

	//correct copy slice
	slice7 := make([]int, len(slice5), len(slice5))
	copy(slice7, slice5)
	fmt.Println(slice7)

	//slice parts
	fmt.Println("part of slice", slice7, slice7[1:5], slice7[:2], slice7[5:])

	//slice from array
	arr := [...]int{5, 6, 7}
	sl8 := arr[:]
	arr[1] = 8
	fmt.Println("slice from array", sl8)
}
