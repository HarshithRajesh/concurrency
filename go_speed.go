package main

import (
	"fmt"
	"time"
)

func printMessage(msg string) {
	for i := 0; i < 3; i++ {
		fmt.Println(msg, i)
		time.Sleep(100 * time.Millisecond)
	}
}

func main() {
	go printMessage("Goroutine")       // runs concurrently
	printMessage("Go is fast")         // runs in main thread
	time.Sleep(500 * time.Millisecond) // wait for goroutine to finish
}
