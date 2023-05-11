package main

import "fmt"

func fibonacci() func(int) int {
	fncSlice := []int{0, 1}
	return func(x int) int {
		if x < 2 {
			return fncSlice[x]
		} else {
			fncSlice = append(fncSlice, (fncSlice[x-1] + fncSlice[x-2]))
			return fncSlice[x]
		}
	}
}

func main() {

	fbf := fibonacci()
	for i := 0; i < 20; i++ {
		fmt.Println(fbf(i))
	}
}
