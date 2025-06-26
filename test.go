// example of simple goroutine hello go func
package main

import (
	"fmt"
	"time"
)

func sayHello() {
	for i := 0; i < 3; i++ {
		fmt.Println("Hello, Go!")
		time.Sleep(500 * time.Millisecond)
	}
}

func main() {
	go sayHello()               // This starts a new goroutine fmt.Println("Main function execution")
	time.Sleep(2 * time.Second) // Wait to let the goroutine finish
}
