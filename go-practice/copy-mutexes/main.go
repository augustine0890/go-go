package main

import (
	"fmt"
	"sync"
	"time"
)

type Container struct {
	sync.Mutex // Add a mutex
	counters   map[string]int
}

func (c *Container) inc(name string) {
	c.Lock() // Locking of the mutex
	defer c.Unlock()
	c.counters[name]++
}

func main() {
	c := Container{counters: map[string]int{"a": 0, "b": 0}}

	doIncrement := func(name string, time int) {
		for i := 0; i < time; i++ {
			c.inc(name)
		}
	}

	doIncrement("a", 100)
	doIncrement("a", 100)

	// Wait for the goroutine to finish
	time.Sleep(500 * time.Millisecond)
	fmt.Println(c.counters)
}
