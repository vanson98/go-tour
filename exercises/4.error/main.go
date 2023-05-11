package main

import "fmt"

type ErrNegativeSqrt float64

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}
	z := float64(1000)
	for i := 0; i <= 10; i++ {
		z -= (z*z - x) / (2 * z)

	}
	return z, nil

}

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("Can't Sqrt negative number %v", float64(e))
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
