package main

import "fmt"

var c chan int

func main() {
	//create channel
	c := make(chan string)
	//start writing goroutine
	go greet(c)
	for i := 0; i < 5; i++ {
		//read 2 lines from channel
		fmt.Println(<-c, ",", <-c)
	}

	stuff := make(chan int, 7)
	for i := 0; i < 19; i = i + 3 {
		stuff <- i
	}
	close(stuff)
	fmt.Println("Res", process(stuff))
}

func greet(c chan<- string) {
	//inf loop
	for {
		//and write 2 lines
		//sub-program will be blocked until someone reads lines
		c <- fmt.Sprintf("Value One")
		c <- fmt.Sprintf("Value Two")
	}
}

func process(input <-chan int) (res int) { //input <- chan int is read-only channel for func
	for r := range input {
		res += r
	}
	return res
}
