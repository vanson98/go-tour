// package main

// import "fmt"

// type Person struct {
// 	Name string
// 	Age  int
// }

// /*
// One of the most ubiquitous interfaces is Stringer defined by the fmt package.
//  type Stringer interface {
//      String() string
//  }

//  A Stringer is a type that can describe itself as a string.
//  The fmt package (and many others) look for this interface to print values.
// */

// func (p Person) String() string {
// 	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
// }

// func main() {
// 	a := Person{"Arthur Dent", 42}
// 	z := Person{"Zaphod Beeblebrox", 9001}
// 	//a.String()
// 	fmt.Println(a, z)
// }

package main

import "fmt"

type IPAddr [4]byte

func (ip IPAddr) String() string {
	return fmt.Sprintf("%v.%v.%v.%v", ip[0], ip[1], ip[2], ip[3])
}

func main() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}
