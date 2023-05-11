package main

import "fmt"

// explicit variable
var c, python, java bool

// implicit variable
// must assign value for variable
var name = "Nguyen "

// variable with initializers
var i, j, l int = 1, 2, 3
var age, handsome, mode = 2, true, "test"

func main() {
	// short variable declaration - implicit variable v2 (only used in function level)
	lastname := "Son"
	fmt.Print(lastname)
}
