package main

import "fmt"

func greeting(myChannel chan string) {
	myChannel <- "hi"
}

func abc(channel chan string) {
	channel <- "a"
	channel <- "b"
	channel <- "c"
}

func def(channel chan string) {
	channel <- "d"
	channel <- "e"
	channel <- "f"
}

func main() {
	// Declare a variable to hold a channel
	// Create the channel
	myChannel := make(chan string)
	go greeting(myChannel)
	receivedValue := <-myChannel
	fmt.Println(receivedValue)
	// fmt.Println(<-myChannel)

	channel1 := make(chan string)
	channel2 := make(chan string)
	go abc(channel1)
	go def(channel2)
	fmt.Println(<-channel1)
	fmt.Println(<-channel2)
	fmt.Println(<-channel1)
	fmt.Println(<-channel2)
	fmt.Println(<-channel1)
	fmt.Println(<-channel2)
}
