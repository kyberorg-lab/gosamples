package main

import (
	"fmt"
	"math/rand"
	"time"
)

func isYearIsLeap(year int) bool {
	return year%400 == 0 || (year%4 == 0 && year%100 != 0)
}

func randomNumber() int {
	return rand.Intn(10) + 1 //func gives nums from 0 to 9
}

func main() {
	fmt.Printf("Current year %d \n", time.Now().Year())
	fmt.Printf("Current weekday %s \n", time.Now().Weekday().String())

	var leapYear = isYearIsLeap(time.Now().Year())
	fmt.Printf("Current year is leap ? %t \n", leapYear)
	fmt.Printf("Random Number is %d\n", randomNumber())
}
