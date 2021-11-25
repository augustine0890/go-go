package main

import "fmt"

func main() {
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	<-ch      // receive value from ch
	close(ch) // close it
	a := <-ch
	b := <-ch
	fmt.Println(a, b)
}
