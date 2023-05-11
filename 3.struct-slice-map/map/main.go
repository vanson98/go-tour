package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

func main() {
	fmt.Println("========= MAP ===========")

	// A map maps keys to values
	// Zero value of a map is nil. A nil map has no keys, nor can keys be added

	var m = make(map[string]Vertex, 0)
	m["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}
	fmt.Println(m["Bell Labs"])

	// Map literal
	var m2 = map[string]Vertex{
		"London": Vertex{
			Lat:  123,
			Long: 456,
		},
		"Ha Noi": {
			222,
			999,
		},
	}

	fmt.Print(m2)
}
