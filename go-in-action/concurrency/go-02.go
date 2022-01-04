package main

import (
	"fmt"
	"runtime"
	"sync"
)

// wg is used to wait for the program to finish
var wg sync.WaitGroup

func main() {
	// Allocate 1 logical memory processors for scheduler to used
	// Change the runtime to allocate a logical processor for every available physical processor.
	runtime.GOMAXPROCS(2)

	wg.Add(2)

	fmt.Println("Create Goroutines")

	go printPrime("A")
	go printPrime("B")

	// Wait for the goroutines to finish
	fmt.Println("Wait for the goroutines to finish")
	wg.Wait()

	fmt.Println("Terminating Program...")
}

// printPrime displays prime numbers for the first 5000 numbers.
func printPrime(prefix string) {
	// Schedule the call to Done to tell main we are done
	defer wg.Done()

next:
	for outer := 2; outer < 5000; outer++ {
		for inner := 2; inner < outer; inner++ {
			if outer%inner == 0 {
				continue next
			}
		}
		fmt.Printf("%s:%d\n", prefix, outer)
	}
	fmt.Println("Completed", prefix)
}
