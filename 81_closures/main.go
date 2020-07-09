package main

import (
	"fmt"
	"time"
)

func main() {
	myTimer := getTimer()
	defer myTimer()

	f := func() {
		myTimer()
	}

	f()

	time.Sleep(1 * time.Second)
}

func getTimer() func() {
	start := time.Now()
	return func() {
		fmt.Printf("Time from start %v", time.Since(start))
	}
}
