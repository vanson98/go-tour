package main

import (
	"fmt"
	"math"
)

type Abser interface {
	Abs() float64
}

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

type Vertex struct {
	X, Y float64
}

// A type implements an interface by implementing its methods. There is no explicit declaration of intent, no "implements" keyword
func (v *Vertex) Abs() float64 {
	if v == nil {
		return 0
	}
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	var a Abser
	f := MyFloat(-math.Sqrt2)
	v := Vertex{6, 3}

	// A value of interface type can hold any value that implements methods signarutes
	a = f
	describe(a)
	a = &v
	describe(a)

	// In the following line, v is a Vertex (not *Vertex) and does not implements abser
	// a = v
	fmt.Println(a.Abs())

	// if the concrete value inside the interface itself is nil, the method will be called with a nil receiver
	// in some languages this would trigger a null pointer exception, but in go it is common to write methods that gracefully handle being called with a nil receiver
	// note that an interface value that holds a nil concrete value is itself non-nil
	var nv *Vertex
	a = nv
	describe(a)
	fmt.Println(a.Abs())

	// nil interface value holds neither value nor concrete type
	// calling a method on a nil interface is a run-time error because there is no type inside the interface typle to indicate which concrete method to call
	var niv Abser
	describe(niv)
	//niv.Abs()

	// the interface type that specifies zero methods is known as the empty interface
	// an empty interface may hold values of any type
	// Empty interfaces are used by code that handles values of unknown type. For example, fmt.Print takes any number of arguments of type interface{}.
	var i interface{}
	describe_empty_interface(i)
	i = 42
	describe_empty_interface(i)

}

func describe(i Abser) {
	// interface value can be thought of as a tuple of a value and a concrete type
	fmt.Printf("(%v, %T)\n", i, i)
}

func describe_empty_interface(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}
