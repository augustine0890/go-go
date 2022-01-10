package main

import (
	"fmt"
	"time"
)

func main() {
	output1 := make(chan string)
	output2 := make(chan string)
	go server1(output1)
	go server2(output2)

	select {
	case s1 := <-output1:
		fmt.Println(s1)
	case s2 := <-output2:
		fmt.Println(s2)
	}
}

// Assume that the func server1 and server2 are in communicationg with 2 such servers.
// The server which responds first is chosen by the select and the other response is ignored.
func server1(ch chan string) {
	time.Sleep(4 * time.Second)
	ch <- "from server 1"
}

func server2(ch chan string) {
	time.Sleep(2 * time.Second)
	ch <- "from server 2"
}
