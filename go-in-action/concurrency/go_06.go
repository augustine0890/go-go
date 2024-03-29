// This sample program demonstrates how to use an unbuffered channel to simulate a game of tennis between two goroutines.
package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// wg is used to wait for the program to finish
var wg sync.WaitGroup

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	// Create an unbuffered channel
	court := make(chan int)

	wg.Add(2)

	// Launch two players
	go player("Nadal", court)
	go player("Djokovic", court)

	// Start the set
	court <- 1

	// Wait for the game to finish
	wg.Wait()
}

// player simulates a person playing the game of tennis
func player(name string, court chan int) {
	// Schedule the call to Done to tell main we are done
	defer wg.Done()

	for {
		// Wait for the ball to be hit back to us.
		ball, ok := <-court
		if !ok {
			// If the channel was closed we won.
			fmt.Printf("Player %s Won\n", name)
			return
		}

		// Pick a random number and see if we miss the ball.
		n := rand.Intn(100)
		if n%11 == 0 {
			fmt.Printf("Player %s missed\n", name)

			// Close the channel to signal we lost
			close(court)
			return
		}

		// Display and then increment the hit count by one.
		time.Sleep(time.Millisecond * 500)
		fmt.Printf("Player %s hit %d\n", name, ball)
		ball++

		// Hit the ball back to the opposing player
		court <- ball
	}
}
