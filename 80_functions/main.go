package main

import "fmt"

var readyFlag = "not ready"

//this func runs before others
func init() {
	readyFlag = "ready"
}

func main() {
	fmt.Println("We are", readyFlag, "to go")
	funcWithNoParams()
	functionWithTwoDifferentParams(7, "ja")
	twoParamsFunc(1, 2)
	functionThatReturnsValue(5)
	multiInput(1, 2, 3, 42)

	slice1 := []int{1, 2, 4, 8}
	//slice will be interpreted as int, int ...
	multiInput(slice1...)

	result, err := multiReturn(1, 2, 42, -42)
	fmt.Println("Result is", result, "Error is", err)
}

func funcWithNoParams() {
	fmt.Println("I am function without params")
}

func twoParamsFunc(a, b int) {
	fmt.Println("1st param is", a, "2rd param is", b)
}

func functionWithTwoDifferentParams(a int, b string) {
	fmt.Println("1st param (int) is", a, "2rd param (string) is", b)
}

func functionThatReturnsValue(num int) int {
	return num * num * 2
}

func multiInput(incomingValues ...int) (res int) {
	for i := range incomingValues {
		res += incomingValues[i]
	}
	return res
}

func multiReturn(values ...int) (int, error) {
	res := 0
	for i := range values {
		if values[i] < 0 {
			return 0, fmt.Errorf("only natural numbers expected here - given: %d", values[i])
		}
		res += values[i]
	}
	return res, nil
}
