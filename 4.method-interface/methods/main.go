package main

import (
	"math"
)

type Vertex struct {
	X, Y float64
}

type AppFloat float64

// method is just a function with a receiver argument
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// normal function
func AbsFunc(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// only declare a method with a receiver whose type is defined in the same package as the method
func (f AppFloat) Abs() float64 {
	return float64(f)
}

// pointer reveiver (modify receiver)
func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func ScaleFunc(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex{3, 4}
	p := &v

	// function that take a value argument must take a value of that specific type
	AbsFunc(v)
	// while methods with value receivers take either a value or a pointer as the receiver when they are called
	v.Abs()
	p.Abs()

	// methods with pointer receivers take either a value or a pointer as the receiver
	p.Scale(2)
	v.Scale(2) // go interprets the statement v --> (&v)

	// func with pointer argument must take a pointer
	ScaleFunc(&v, 2)

}
