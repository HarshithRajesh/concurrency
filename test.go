// example of simple goroutine hello go func
package main

import (
	"fmt"
	"sync"
	"time"
)

func sayHello() {
	for i := 0; i < 3; i++ {
		fmt.Println("Hello, Go!")
		time.Sleep(500 * time.Millisecond)
	}
}

func main() {
	// go sayHello()               // This starts a new goroutine fmt.Println("Main function execution")
	// time.Sleep(2 * time.Second) // Wait to let the goroutine finish
	// var wg sync.WaitGroup
	// wg.Add(1)
	// go func() {
	// 	defer wg.Done()
	// 	fmt.Println("1st goroutine sleeping...")
	// 	time.Sleep(1)
	// }()
	// wg.Add(1)
	// go func() {
	// 	defer wg.Done()
	// 	fmt.Println("2nd goroutine sleeping...")
	// 	time.Sleep(2)
	// }()
	// wg.Wait()
	// fmt.Println("All the goroutine are sleeping")
	//
	hello := func(wg *sync.WaitGroup, id int) {
		defer wg.Done()
		fmt.Printf("Hello from %v!\n", id)
	}
	const numGreeters = 5
	var wg sync.WaitGroup
	wg.Add(numGreeters)
	for i := 0; i < numGreeters; i++ {
		go hello(&wg, i+1)
	}
	wg.Wait()
}
