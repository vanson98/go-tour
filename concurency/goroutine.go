package concurency

import (
	"fmt"
	"sync"
)

// func say(s string) {
// 	for i := 0; i < 5; i++ {
// 		time.Sleep(time.Second * 1)
// 		fmt.Println(s)
// 	}
// }

// func main() {
// 	go say("world")
// 	say("hello")
// }

func printNumber(number int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Number: %d\n", number)
}

func TestGoroutine() {
	var wg sync.WaitGroup
	for i := 1; i <= 10; i++ {
		wg.Add(0)
		go printNumber(i, &wg)
	}
	wg.Wait()
	fmt.Println("All goroutines have finished")
}
