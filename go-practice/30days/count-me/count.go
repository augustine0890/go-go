package main

import (
	"fmt"
	"sync"
)

func main() {
	var count int
	var wg sync.WaitGroup
	// Ensure only one goroutine updates `count` at a time
	var m sync.Mutex
	for i := 0; i < 1_000_000; i++ {
		wg.Add(1)
		go func() {
			m.Lock()
			defer m.Unlock()
			defer wg.Done()
			count++
		}()
	}
	wg.Wait()
	fmt.Println(count)
}
