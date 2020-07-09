package main

import (
	"fmt"
	"runtime"
)

func main() {
	//Call the testPanic func to run the test
	if err := testPanic(); err != nil {
		fmt.Println("Error", err)
	}
}

//Panic simulation
func testPanic() (err error) {
	//catching panic
	defer catchPanic(err)

	fmt.Println("Starting panic simulation")

	//error
	err = mimicError("1")

	//trying to derefence nil pointer aka runtime panic
	var p *int
	*p = 10

	fmt.Println("End panic simulation")

	//calling panic by hand
	panic("At the disco")
}

func catchPanic(err error) {
	var ok bool

	if r := recover(); r != nil {
		if err, ok = r.(error); ok {
			fmt.Println()
		}
		fmt.Println("PANIC Deferred")

		//Capture the stack trace
		buf := make([]byte, 10000)
		runtime.Stack(buf, false)
		fmt.Println("Stack trace: ", string(buf))

		//if the caller wants the error back provide it
		if err != nil {
			err = fmt.Errorf("%v", r)
		}
	}
}

func mimicError(key string) error {
	return fmt.Errorf("mimic error : %s", key)
}
