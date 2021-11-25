package main

import "fmt"

// ch with capacity of 1 (will get deadlock if want to recive more than 1)
func main() {
	ch := make(chan int, 1)
	ch <- 1
	a, ok := <-ch
	fmt.Println(a, ok) // 1, true
	close(ch)
	b, ok := <-ch
	fmt.Println(b, ok) // 0, false
}
