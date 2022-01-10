// This sample program demonstrates how to use an unbuffered channel to simulate a replay race between four goroutines.
package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	// Create an unbuffered channel
	baton := make(chan int)

	// Add a count of one for the last runner
	wg.Add(1)

	// First runner to his mark
	go Runner(baton)

	// Start the race
	baton <- 1

	wg.Wait()
}

// Runner simulate a person running in the relay race
func Runner(baton chan int) {
	var newRunner int

	// Wait to receive the baton
	runner := <-baton

	// Start running around the track
	fmt.Printf("Runner %d running with baton\n", runner)

	// New runner to the line
	if runner != 4 {
		newRunner = runner + 1
		fmt.Printf("Runner %d to the line\n", newRunner)
		go Runner(baton)
	}

	// Running around the track
	time.Sleep(250 * time.Millisecond)

	// Is the race over
	if runner == 4 {
		fmt.Printf("Runner %d finished. Race over\n", runner)
		wg.Done()
		return
	}

	// Exchange the baton for the next runner.
	fmt.Printf("Runner %d exchange with runner %d\n", runner, newRunner)

	baton <- newRunner
}
