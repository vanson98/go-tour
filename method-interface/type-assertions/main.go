package main

import "fmt"

type Vertex struct {
	X float64
	Y float64
}

func main() {
	var si interface{} = "hello"
	var vi interface{} = Vertex{1, 4}

	// type assertion provides access to an interface value's underlying concrete value
	s := si.(string)
	x := vi.(Vertex)
	fmt.Println(s, x)

	// to test whether an interface value holds a specific type, a type assertion can return two values:
	// the underlying value and a boolean value that report whether the assertion succeeded
	// if si holds a string, then s will be the underlying value and ok will be true
	// if not, ok will be false and s will be the zero value of type string, and no panic occurs
	s, ok := si.(string)
	fmt.Println(s, ok)

	f, ok := si.(float64)
	fmt.Println(f, ok)

	// if si does not hold a float value, the statement will trigger a panic
	f = si.(float64) // panic
	fmt.Println(f)

}
