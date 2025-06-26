package main

import "fmt"

func sum(a, b int, result chan int) {
	result <- a + b
}

func main() {
	result := make(chan int)
	go sum(3, 4, result)
	fmt.Println("Sum:", <-result) // Receives the value from the channel
}
