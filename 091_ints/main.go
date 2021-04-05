package main

import (
	"fmt"
	"sort"
)

type MyStruct struct {
	Num  int
	Name string
}

type MyInt int
type withFiles bool

//aliasing
var WithFiles = withFiles(true)
var WithoutFiles = withFiles(false)

func (m MyInt) showYourSelf() {
	fmt.Printf("%T %v\n", m, m)
}

func (m *MyInt) add(i MyInt) {
	*m = *m + MyInt(i)
}

type mySliceStruct []MyStruct

func (m mySliceStruct) Less(i int) bool
func (m mySliceStruct) Len() int
func (m mySliceStruct) Swap(i, j int)

func main() {
	myVar := MyStruct{
		Num:  0,
		Name: "Void",
	}

	println(myVar)

}
