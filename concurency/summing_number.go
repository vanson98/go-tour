package concurency

import (
	"fmt"
	"sync"
)

func sum(numbers []int, wg *sync.WaitGroup, result *int, mux *sync.Mutex) {
	defer wg.Done()
	localSum := 0
	for _, num := range numbers {
		localSum += num
	}
	mux.Lock()
	*result += localSum
	mux.Unlock()
}

func TestSumming() {
	numbers := []int{1, 4, -2, 3, 4, 1, 7, 9, 6, 5, 8, 5}
	chunkSize := 2
	totalSum := 0
	var wg sync.WaitGroup
	var mux sync.Mutex

	for i := 0; i < len(numbers); i += chunkSize {
		end := i + chunkSize
		if end > len(numbers) {
			end = len(numbers)
		}
		wg.Add(1)
		go sum(numbers[i:end], &wg, &totalSum, &mux)
	}
	wg.Wait()
	fmt.Printf("Total sum: %d\n", totalSum)
}
