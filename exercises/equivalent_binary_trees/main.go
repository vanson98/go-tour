package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

func main() {
	mytree1 := tree.New(1)
	mytree2 := tree.New(2)
	Same(mytree1, mytree2)
}

func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)
	x := make([]int, 10)
	y := make([]int, 10)
	go Walk(t1, ch1)
	for val := range ch1 {
		x = append(x, val)
	}

	go Walk(t2, ch2)
	for val := range ch2 {
		y = append(y, val)
	}

	// for {
	// 	select {
	// 	case val1 := <-ch1:
	// 		x = append(x, val1)
	// 	case val2 := <-ch2:
	// 		y = append(y, val2)
	// 	default:
	// 		break
	// 	}

	// }
	fmt.Println(x)
	fmt.Println(y)
	return checkSliceEqual(x, y)
}

func checkSliceEqual(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

var i int = 1

func Walk(t *tree.Tree, ch chan int) {

	if t.Left != nil {
		Walk(t.Left, ch)
	}

	// if i == 10 {
	// 	ch <- t.Value
	// 	println("end")
	// 	close(ch)
	// } else {
	//	i = i + 1
	ch <- t.Value
	//}
	if t.Right != nil {
		Walk(t.Right, ch)
	}

}

// left - root - right
func inOrderTraverse(t *tree.Tree) {
	if t.Left != nil {
		inOrderTraverse(t.Left)
	}
	fmt.Println(t.Value)
	if t.Right != nil {
		inOrderTraverse(t.Right)
	}

}

// left - right - root
func postOrderTraverse(t *tree.Tree) {
	if t.Left != nil {
		postOrderTraverse(t.Left)
	}

	if t.Right != nil {
		postOrderTraverse(t.Right)
	}
	fmt.Println(t.Value)
}

// root - left - right
func preOrderTraverse(t *tree.Tree) {
	fmt.Println(t.Value)
	if t.Left != nil {
		preOrderTraverse(t.Left)
	}
	if t.Right != nil {
		preOrderTraverse(t.Right)
	}
}
