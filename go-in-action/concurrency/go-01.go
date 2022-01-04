// This sample program demonstrates how to create goroutines and how the scheduler dehaves.
package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	// Allocate 1 logical processor for the scheduler to use
	// fmt.Println("CPU Num:", runtime.NumCPU()) // number of physical processors that are available

	runtime.GOMAXPROCS(2)

	// wg is used to wait for the program to finish
	// Add a count of two, one for each goroutine
	var wg sync.WaitGroup
	wg.Add(2)

	fmt.Println("Start Goroutine")

	// Declare an anonymous function and create a goroutine
	go func() {
		// Schedule the call to Done to tell main we are done.
		defer wg.Done()

		// Display the alphabet three times
		for count := 0; count < 3; count++ {
			for char := 'a'; char <= 'z'; char++ {
				fmt.Printf("%c ", char)
			}
		}
	}()

	go func() {
		// Guarantee to be called once per goroutine is finished.
		defer wg.Done()
		for count := 0; count < 3; count++ {
			for char := 'A'; char <= 'Z'; char++ {
				fmt.Printf("%c ", char)
			}
		}
	}()

	// Wait for the goroutine to finish
	fmt.Println("Waiting To Finish")
	wg.Wait()

	fmt.Println("\nTerminating Program")
}
