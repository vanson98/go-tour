package main

import "fmt"

type chanWriter struct {
	ch chan byte
}

func newChanWriter() *chanWriter {
	return &chanWriter{make(chan byte, 1024)}
}

func (w *chanWriter) Chan() <-chan byte {
	return w.ch
}

func (w *chanWriter) Write(p []byte) (int, error) {
	n := 0
	for _, b := range p {
		w.ch <- b
		n++
	}
	return n, nil
}

func (w *chanWriter) Close() error {
	close(w.ch)
	return nil
}

func main() {
	writer := newChanWriter()
	go func() {
		defer writer.Close()
		writer.Write([]byte("Stream "))
		writer.Write([]byte("me!"))
	}()
	for c := range writer.Chan() {
		fmt.Printf("%c", c)
	}
	fmt.Println()
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
