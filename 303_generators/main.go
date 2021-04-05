package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// func returns channel
	c := boring("boring!")
	for i := 0; i < 5; i++ {
		//read from channel
		fmt.Printf("You say: %q\n", <-c)
	}
	fmt.Println("You are boring, I am leaving.")
}

func boring(msg string) <-chan string { //return read-only channel
	c := make(chan string)
	go func() {
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
		}
	}()
	return c
}
