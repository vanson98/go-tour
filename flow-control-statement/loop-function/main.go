package main

import "fmt"

func Sqrt(x float64) {
	z := float64(1000)
	for i := 0; i <= 10; i++ {
		z -= (z*z - x) / (2 * z)
		fmt.Println(z)
	}

}

func main() {
	Sqrt(9)
}
