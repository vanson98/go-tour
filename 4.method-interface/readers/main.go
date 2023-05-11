package main

import (
	"fmt"
	"io"
	"strings"
)

func main() {
	r := strings.NewReader("Hello, Reader!")
	b := make([]byte, 8)

	for {
		// the io package specifies the io.Reader interface, which represents the read end of a stream of data
		// The Go standard library contains many implementations of this interface, including files, network connection, compressors, ciphers, and others
		// The io.Reader interface has a Read method
		// func (T) Read(b []byte) (n int, err error)
		// Read populates the given byte slice with data and returns the number of bytes populated and an error value. It returns an
		// io.Reader interface has a Read method
		n, err := r.Read(b)
		// Read populates the given byte slice with data and returns the number of bytes populated and an error value. It returns and io.EOF error
		// when the stream ends
		fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
		fmt.Printf("b[:n] = %q\n", b[:n])
		if err == io.EOF {
			break
		}
	}
}
