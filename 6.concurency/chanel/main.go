package main

// ======================== Range and Close ======================

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
