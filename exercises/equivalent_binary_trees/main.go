package main

import "golang.org/x/tour/tree"

func Walk(t *tree.Tree, ch chan int) {
	ch <- t.Value
}

func main() {
	mytree := tree.New(1)
	println(mytree)
}
