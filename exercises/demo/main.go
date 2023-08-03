package main

type sonWriter struct {
	ch chan byte
}

func main() {

}

// ==============================================================================================
// package main
// import (
// 	"bytes"
// 	"fmt"
// 	"os"
// )

// func main() {
// 	// Writer
// 	investment := []string{
// 		"tcb",
// 		"vic",
// 		"tar",
// 	}

// 	var writer bytes.Buffer

// 	for i := 0; i < len(investment); i++ {
// 		n, err := writer.Write([]byte(investment[i]))

// 		if err != nil {
// 			fmt.Println(err)
// 			os.Exit(1)
// 		}

// 		if n != len(investment[i]) {
// 			fmt.Println("Fail to write")
// 			os.Exit(1)
// 		}
// 	}

// 	fmt.Println(writer.String())
// }
