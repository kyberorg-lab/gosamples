package main

import (
	"fmt"
	"math"
	"math/cmplx"
	"math/rand"
	"runtime"
	"time"
)

const (
	// Create a huge number by shifting a 1 bit left 100 places.
	// In other words, the binary number that is 1 followed by 100 zeroes.
	Big = 1 << 100
	// Shift it right again 99 places, so we end up with 1<<1, or 2.
	Small = Big >> 99
)

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func typeShow() {
	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)
}

func needInt(x int) int { return x*10 + 1 }
func needFloat(x float64) float64 {
	return x * 0.1
}

func hugeNums() {
	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))
	//Big is too huge for int
	//fmt.Println(needInt(Big))
}

func circleShort(init, max int) int {
	sum := init
	for sum < max {
		sum += sum
	}
	return sum
}

func os() string {
	switch os := runtime.GOOS; os {
	case "darwin":
		return "OS X"
	case "linux":
		return "Linux"
	case "windows":
		return "win: change this shit"
	default:
		return os
	}
}

func whenSaturday() {
	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}
}

//goland:noinspection ALL
func deferExample() {
	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")
}

func main() {
	//random with seed
	rand.Seed(time.Now().UnixNano())
	fmt.Println("Rand number is", rand.Intn(10))
	//math
	fmt.Printf("Now you have %g problems.\n", math.Sqrt(7))
	//Pi
	fmt.Println(math.Pi)
	//split
	fmt.Println(split(17))
	//types
	typeShow()
	//huge
	hugeNums()
	//short syntax for 'for'
	fmt.Println(circleShort(1, 1000))
	//my os
	fmt.Println("Go runs at", os())
	//Saturday
	whenSaturday()
	//defer example
	deferExample()
}
