package main

import "fmt"

// A struct is a collection of fields
type Vertex struct {
	X int
	Y int
}

type Car struct {
	A string
	B bool
}

func main() {
	v := Vertex{1, 2}
	fmt.Println(v)

	// struct fields are accessed using dot
	v.X = 5

	// struct fields can be accessed throught a struct pointer
	p := &v
	p.X = 1e9 // = (*p).X

	// struct literal denotes a newly allocated struct value by listing the values of its fields
	var v1 = Vertex{1, 2}
	var v2 = &Vertex{1, 2}

	fmt.Println(v1, v2)
	fmt.Println(v)
	fmt.Println(Car{"Toyota Camry", true})
}
