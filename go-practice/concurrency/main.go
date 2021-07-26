package main

import (
	"fmt"
	"time"
)

/*
Concurrency is the ability of different parts or units of a program, or problem to be executed
out-of-order or in partial order, without affecting the final outcome.
Parallelism refer to tasks that are executed simultaneously (at the same time)
*/

func printNumber() {
	i := 0
	for i < 3 {
		time.Sleep(1 * time.Second)
		i++
		fmt.Println(i)
	}
}

func main() {
	// panic("Show me the goroutine")
	fmt.Println("launch goroutine")
	go printNumber()
	fmt.Println("launch goroutine")
	go printNumber()
	time.Sleep(5 * time.Second)

	var received int
	ch1 := make(chan int, 2)
	ch1 <- 1
	ch1 <- 2
	close(ch1)
	received, ok := <-ch1
	fmt.Println(received, ok)
}
