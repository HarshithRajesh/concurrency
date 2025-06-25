package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var data int
	var memoryAccess sync.Mutex
	go func() {
		memoryAccess.Lock()
		data++
		memoryAccess.Unlock()
	}()
	memoryAccess.Lock()
	time.Sleep(1 * time.Second)
	if data == 0 {
		fmt.Printf("the value is %v.\n", data)
	} else {
		fmt.Printf("the value is %v.\n", data)
	}
	memoryAccess.Unlock()
}
