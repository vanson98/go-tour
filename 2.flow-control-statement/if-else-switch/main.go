package main

import (
	"fmt"
	"math"
	"runtime"
	"time"
)

func main() {
	sum := 0
	// init statement ; condition expresion ; post statement
	for i := 0; i < 10; i++ {
		sum += i
	}

	// The init post statement are optional  == while loop
	for sum < 1000 {
		sum += sum
	}
	fmt.Println((sum))

	// infinite loop
	// for {
	// 	fmt.Println("H")
	// }

	// if statement
	if sum > 100 {
		fmt.Print("sun is larger than 100")
	}

	// go only runs the selected case, not all the cases that follow
	// break statement automatically in Go
	// Go's switch cases need not be constants, and the values involved need not be integers
	// Switch cases evaluate cases from top to bottom, stopping when a case succeeds
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux")
	default:
		fmt.Printf("%s /n", os)
	}

	// switch with no condition == switch true
	// clean way to write long if else chains
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon!")
	default:
		fmt.Println("Good evening")
	}
}

func pow(x, n, lim float64) float64 {
	// variable declared by the statement are only in scope until the end of the if  or inside any of the else blocks
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Print(v)
	}
	return lim
}
