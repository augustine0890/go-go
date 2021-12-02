package main

import "fmt"

// Deferred funcs are executed in LIFO order
func main() {
	for i := 0; i < 5; i++ {
		defer fmt.Printf("%d\n", i)
	}
}
