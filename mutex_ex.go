package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	mu    sync.Mutex
	value int
}

func (c *Counter) increment() {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.value++
}

func main() {
	// var counter Counter
	// var wg sync.WaitGroup
	//
	// for i := 0; i < 1000; i++ {
	// 	wg.Add(1)
	// 	go func() {
	// 		defer wg.Done()
	// 		counter.increment()
	// 	}()
	// }
	//
	// wg.Wait()
	// fmt.Println("Final counter value:", counter.value)
	var count int
	var lock sync.Mutex
	increment := func() {
		lock.Lock()
		defer lock.Unlock()
		count++
		fmt.Printf("Incrementing: %d\n", count)
	}
	decrement := func() {
		lock.Lock()
		defer lock.Unlock()
		count--
		fmt.Printf("Decrementing: %d\n", count)
	}
	// Increment
	var arithmetic sync.WaitGroup
	for i := 0; i <= 5; i++ {
		arithmetic.Add(1)
		go func() {
			defer arithmetic.Done()
			increment()
		}()
		for i := 0; i <= 5; i++ {
			arithmetic.Add(1)
			go func() {
				defer arithmetic.Done()
				decrement()
			}()
		}
		arithmetic.Wait()
		fmt.Println("Arithmetic complete.")
	}
}
