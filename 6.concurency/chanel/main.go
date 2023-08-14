package main

// ========================  RANDOM Select statement ===================

// Go program to illustrate the
// concept of select statement

import "fmt"

// function 1
func portal1(channel1 chan string) {
	for i := 0; i <= 3; i++ {
		channel1 <- "Welcome to channel 1"
	}

}

// function 2
func portal2(channel2 chan string) {
	channel2 <- "Welcome to channel 2"
}

// main function
func main() {

	// Creating channels
	R1 := make(chan string)
	R2 := make(chan string)

	// calling function 1 and
	// function 2 in goroutine
	go portal1(R1)
	go portal2(R2)

	// the choice of selection
	// of case is random
	select {
	case op1 := <-R1:
		fmt.Println(op1)
	case op2 := <-R2:
		fmt.Println(op2)
	}
}

// =========================================== DEMO SELECT 2=========================================
// The select statement lets a goroutine wait on multiple communication operations.
// A select blocks until one of its cases can run, then it executes that case. It chooses one at random if multiple are ready.

// import "fmt"

// func fibonacci(c, quit chan int) {
// 	x, y := 0, 1
// 	for {
// 		select {
// 		case c <- x:
// 			x, y = y, x+y
// 		case <-quit:
// 			fmt.Println("quit")
// 			return
// 		}
// 	}
// }

// func main() {
// 	c := make(chan int)
// 	quit := make(chan int)
// 	go func() {
// 		for i := 0; i < 10; i++ {
// 			fmt.Println(<-c)
// 		}
// 		quit <- 0
// 	}()
// 	fibonacci(c, quit)
// }

// ========================================= DEMO SELECT 1 ====================================
// import (
// 	"fmt"
// 	"time"
// )

// // function 1
// func portal1(channel1 chan string) {
// 	time.Sleep(99 * time.Second)
// 	channel1 <- "Welcome to channel 1"
// }

// // function 2
// func portal2(channel2 chan string) {
// 	time.Sleep(1 * time.Second)
// 	channel2 <- "Welcome to channel 2"
// }

// // main function
// func main() {

// 	// Creating channels
// 	R1 := make(chan string)
// 	R2 := make(chan string)

// 	// calling function 1 and
// 	// function 2 in goroutine
// 	go portal1(R1)
// 	go portal2(R2)

// 	select {

// 	// case 1 for portal 1
// 	case op1 := <-R1:
// 		fmt.Println(op1)

// 	// case 2 for portal 2
// 	case op2 := <-R2:
// 		fmt.Println(op2)
// 	}
// }

// ======================== Range and Close ======================
// A sender can close a channel to indicate that no more values will be sent. Receivers can test whether a channel has been closed by assigning a second parameter to the receive expression: after
// v, ok := <-ch
// ok is false if there are no more values to receive and the channel is closed.

// func fibonacci(n int, c chan int) {
// 	x, y := 0, 1
// 	for i := 0; i < n; i++ {
// 		c <- x
// 		x, y = y, x+y

// 	}
// 	close(c)
// }
// The loop for i := range c receives values from the channel repeatedly until it is closed.
// func main() {
// 	c := make(chan int, 10)
// 	go fibonacci(cap(c), c)
// 	for i := range c {
// 		fmt.Println(i)
// 	}
// }

// Note: Only the sender should close a channel, never the receiver. Sending on a closed channel will cause a panic.
// Another note: Channels aren't like files; you don't usually need to close them. Closing is only necessary when the receiver must be told there are no more values coming, such as to terminate a range loop.

// ========================= Buffered channel ======================
// type Investment struct {
// 	Code string
// }

// func main() {
// 	// Channel can be a buffered. Provide the buffer length as the second argument to "make"
// 	// to initialize buffer channel
// 	ch := make(chan Investment, 2)

// 	// Send to buffer channel block only when buffer is full
// 	ch <- Investment{
// 		Code: "NVL",
// 	}

// 	ch <- Investment{
// 		Code: "VIC",
// 	}

// 	// Receive block when the buffer is empty
// 	fmt.Println(<-ch)
// 	fmt.Println(<-ch)
// }

// ========================= Channel ===============================
// Channels are the typed conduit (mang) through which you can send and receive values with the channel operator, <-
// The data flows in the direction of the arrow
// Like map and slice, channels must be created before use
// By default, sends and receives block until the other side is ready. This allows goroutines to synchronize without explicit locks or condition variables.

// func sum(s []int, c chan int) {
// 	sum := 0
// 	for _, v := range s {
// 		sum += v
// 	}
// 	c <- sum // send sum to c
// }

// func main() {
// 	s := []int{7, 2, 8, -9, 4, 0}

// 	c := make(chan int)
// 	go sum(s[:len(s)/2], c)
// 	go sum(s[len(s)/2:], c)
// 	x, y := <-c, <-c // receive from c

// 	fmt.Println(x, y, x+y)
// }
