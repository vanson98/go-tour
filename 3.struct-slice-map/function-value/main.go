package main

import (
	"fmt"
	"math"
)

func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

func main() {
	// function are values too. They can be passed around just like other values
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}

	//Function values may be used as function arguments and return values.
	fmt.Println(hypot(5, 12))

	fmt.Println(compute(hypot))
	fmt.Println(compute(math.Pow))
}
