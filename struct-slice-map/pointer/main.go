package main

import "fmt"

func main() {
	i, j := 42, 2701

	// pointer holds the memory address of a value
	// the type *T is a pointer to a T value. Its zero value is nil
	// var age *int

	// the & operator generates a pointer to its operand
	p := &i         // point to i
	fmt.Println(*p) // read i through the pointer

	// the * operator denotes the pointer's underlying value
	*p = 21 // set i through the pointer
	fmt.Println(i)

	p = &j       // point to j
	*p = *p / 37 // divide j thought the pointer
	fmt.Println(j)

	// Unlike C, Go has no pointer arithmetic
}
