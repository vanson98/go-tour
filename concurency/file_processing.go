package concurency

import (
	"bufio"
	"fmt"
	"os"
	"sync"
)

func countLines(fileName string, wg *sync.WaitGroup, result *int, mux *sync.Mutex) {
	defer wg.Done()

	file, err := os.Open(fmt.Sprintf("/Users/vanson/Documents/TechNote/%s", fileName))
	if err != nil {
		fmt.Printf("Failed to open file %s: %s\n", fileName, err)
		return
	}
	defer file.Close()

	lineCount := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lineCount++
	}

	mux.Lock()
	*result += lineCount
	mux.Unlock()
}

func TestReadFileConcurency() {
	fileNames := []string{"autotest.txt", "mongodb.txt", "GL.txt", "LinuxDocs.txt", "AngularNote.txt"}
	totalLines := 0
	var wg sync.WaitGroup
	var mux sync.Mutex

	for _, fileName := range fileNames {
		wg.Add(2)
		countLines(fileName, &wg, &totalLines, &mux)
	}

	wg.Wait()
	fmt.Printf("Total number of lines: %d\n", totalLines)
}
