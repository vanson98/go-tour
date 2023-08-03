package main

import (
	"fmt"
	"time"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(time.Second * 1)
		fmt.Println(s)
	}
}

func main() {
	go say("world")
	say("hello")
}
