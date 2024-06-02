package main

import (
	"fmt"
	"io"
)

type MyReader struct {
}

func (mr MyReader) InitMyReader() io.Reader {
	return MyReader{}
}

func (mr MyReader) Read(b []byte) (n int, e error) {
	// for x := range b {
	//  b[x] = byte('A')
	// }
	// return len(b), nil
	b[0] = 'A'
	return 1, nil
}

func main() {
	mr := MyReader{}
	bs := make([]byte, 10)
	n, err := mr.Read(bs)
	fmt.Printf("n=%v err=%v b=%v\n", n, err, bs)
}
