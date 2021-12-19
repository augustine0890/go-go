package main

import (
	"fmt"
	"time"
)

func reportNap(name string, delay int) {
	for i := 0; i < delay; i++ {
		fmt.Println(name, "sleeping")
		time.Sleep(1 * time.Second)
	}
	fmt.Println(name, "wake up!")
}

func send(channel chan string) {
	reportNap("sending goroutine", 2)
	fmt.Println("***sending value***")
	channel <- "a"
	fmt.Println("***sending value***")
	channel <- "b"
}

func main() {
	channel := make(chan string)
	go send(channel)
	reportNap("receiving goroutine", 5)
	fmt.Println(<-channel)
	fmt.Println(<-channel)
}
