package main

import (
	"fmt"
)

const Pi = 3.14

const (
	Big   = 1 << 100
	Small = Big >> 99
)

func needInt(x int) int {
	return x*10 + 1
}

func needFloat(x float64) float64 {
	fmt.Printf("Big is of type %T\n", x)
	return x * 0.1
}

func main() {
	const World string = "Nguyen"
	const Truth = true
	fmt.Println("Go rules", Truth)

	fmt.Printf("Big is type of %T\n", Small)

	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))
}
