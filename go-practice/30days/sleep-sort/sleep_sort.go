package main

import (
	"fmt"
	"sync"
	"time"
)

// goroutines (and deferred functions) are written as function invocation with parameters
func main() {
	var wg sync.WaitGroup
	for _, n := range []int{3, 1, 2} {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			time.Sleep(time.Duration(i) * time.Millisecond)
			fmt.Printf("%d\n", i)
		}(n)
	}
	wg.Wait()
	fmt.Println()
}
