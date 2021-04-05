package main

import "fmt"

func main() {
	fmt.Println("Start")
	//launch func
	go process(0)
	//launch anon func
	go func() {
		fmt.Println("Anon func in goroutine")
	}()

	//several goroutine
	for i := 0; i < 100; i++ {
		go process(i)
	}
	//if nothing below (all goroutines are killed after runtime reach this line)
	//to wait until "several goroutine" stuff finish (makes all interactions) some code below must be added
	fmt.Scanln()
}

func process(i int) {
	fmt.Println("tackle", i)
}
