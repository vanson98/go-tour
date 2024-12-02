package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"
)

func readFile() {
	file, err := os.Open("todolist.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	var buffer []byte = make([]byte, 0)
	onebyte := make([]byte, 1)
	for {
		_, err = file.Read(onebyte)
		if err == io.EOF {
			defer fmt.Println("end of file!")
			break
		} else if err != nil {
			fmt.Println(err)
			break
		}
		buffer = append(buffer, onebyte[0])
	}
	fmt.Println(string(buffer))
}

func readString() {
	r := strings.NewReader("Hello, Reader!")
	b := make([]byte, 8)
	for {
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

// ================================= TEST WRITER ==============================
func writeStringToBuffer() {
	myBuffer := bytes.Buffer{}
	n, err := myBuffer.Write([]byte("Nguyen Van Son Dep Trai"))
	//n, err = fmt.Fprint(&myBuffer, struct{ a int }{a: 1})
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Printf("Number of bytes are written to buffer %d\n", n)
	fmt.Println(string(myBuffer.Bytes()))
}

type myWriter struct {
	buffer []byte
	size   int
}

func (m *myWriter) Write(p []byte) (int, error) {
	if len(p) > m.size {
		return 0, fmt.Errorf("exceed the size of myWriter")
	}
	m.buffer = append(m.buffer, p...)
	return len(m.buffer), nil
}

func testMyWriter() {
	mw := myWriter{buffer: []byte{}, size: 500}
	n, err := fmt.Fprintf(&mw, "Test writer: %s", []byte("Reader and writer"))
	//n, err := mw.Write([]byte("Reader and writer"))
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Number of bytes are written to buffer %d\n", n)
	fmt.Println(string(mw.buffer))
}

func main() {
	//readString()
	//readFile()
	//writeStringToBuffer()
	testMyWriter()
}
