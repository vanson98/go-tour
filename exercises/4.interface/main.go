package main

import "fmt"

type printer interface {
	print()
}

type person struct {
	name string
	age  int
}

func (p *person) print() {
	fmt.Printf("Person Name: %v\n", p.name)
	fmt.Printf("Person Age: %v\n", p.age)
}

func main() {
	var pt printer
	p := person{
		name: "Nguyen Van Son",
		age:  25,
	}
	pt = &p
	pt.print()
}
