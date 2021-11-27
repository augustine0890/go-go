package main

import "fmt"

// Call function after the function complete
// open, close, connect and disconnect, or lock and unlock
// Deferred funcs are run even if a runtime panic occurs
func first() {
	fmt.Println("First")
}

func second() {
	fmt.Println("Second")
}

func main() {
	defer second()
	first()
}
