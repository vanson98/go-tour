package main

import "fmt"

func main() {
	// defer statement defers the execution of a function until the surrounding function returns
	defer fmt.Println("world")

	fmt.Println("hello")

	// deferred calls are executed in last in first out order
	for i := 0; i < 100; i++ {
		defer fmt.Println(i)
	}

}
