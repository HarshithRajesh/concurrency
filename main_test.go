package main

import (
	"sync"
	"testing"
)

func worker(wg *sync.WaitGroup) {
	defer wg.Done()
	// Simulate work
}

func BenchmarkGoroutine(b *testing.B) {
	for n := 0; n < b.N; n++ {
		var wg sync.WaitGroup
		wg.Add(1)
		go worker(&wg)
		wg.Wait()
	}
}
