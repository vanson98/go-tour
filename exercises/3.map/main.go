package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter your string:")
	fmt.Print("->")
	text, _ := reader.ReadString('\n')

	words := strings.Fields(text)
	fmt.Println(words)
	mapResult := make(map[string]int)
	for _, value := range words {
		_, ok := mapResult[value]
		if ok {
			mapResult[value]++
		} else {
			mapResult[value] = 1
		}

	}
	fmt.Println(mapResult)
}
