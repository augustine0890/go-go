package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string)
	go process(ch)
	for {
		time.Sleep(1000 * time.Millisecond)
		select {
		case v := <-ch:
			fmt.Println("received", v)
			return
		default:
			fmt.Println("no value received")
		}
	}
}

func process(ch chan string) {
	time.Sleep(11 * time.Second)
	ch <- "process successful"
}
