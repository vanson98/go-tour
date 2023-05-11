package main

import "fmt"

func main() {
	primes := [6]int{2, 3, 4, 6, 8, 9}
	// A slice is a dynamically-size, flexible view into the elements of an array
	// This selects a half-open range which includes the first element, but excludes the last one
	var s []int = primes[3:6]

	fmt.Println(primes)
	fmt.Println(s)

	fmt.Println("========== Slices are like references to array ===========")

	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}

	fmt.Println(names)

	// Slice does not store any data, it just describes a section of an underlying array
	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b)

	// Changing the elements of a slice modifies the corresponding elements of its underlying array
	b[0] = "XXX"
	fmt.Println(a, b)
	fmt.Println(names)

	fmt.Println("========== Slices literals ===========")

	// slice literal is like an array literal without the length
	// This create a array as bellow, then builds a slice that references it
	q := []int{2, 3, 5, 7, 23, 66}
	fmt.Println(q)

	r := []bool{true, false, true, true, false, true}
	fmt.Println(r)

	st := []struct {
		i int
		b bool
	}{
		{2, true},
		{23, false},
	}
	fmt.Println(st)

	fmt.Println("========== A Slice Defaults ===========")

	// The default is zero for the low bound and the length of the slice for the high bound
	q = q[:5]
	fmt.Println(q)
	q = q[3:]
	fmt.Println(q)

	fmt.Println("=========== Creating a slice with make ===========")

	// Slice can be created with built-in make function
	// The make function allocates a zeroed array and returns a slice that refers to that array

	//a1 := make([]int, 5)  // len(a) = 0

	b1 := make([]int, 5, 10) // len(b) = 5, cap(b) = 10

	fmt.Printf("%s leng=%d cap=%d %v\n", b1, len(b1), cap(b1), b1)

}
