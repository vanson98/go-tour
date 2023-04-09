package main

import "fmt"

func add(x, y int) int {
	return x + y
}

// mutiple results
func swap(x, y string) (string, string) {
	return y, x
}

// naked return
func split(sum int) (a, b int) {
	a = sum * 2
	b = sum - 10
	return
}

func main() {
	fmt.Println(add(42, 13))
	a, b := swap("Nguyen", "Son")
	fmt.Println(a, b)
	fmt.Println(split(100))
}
