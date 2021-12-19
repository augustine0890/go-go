package main

import "fmt"

func greeting(myChannel chan string) {
	myChannel <- "hi"
}

func main() {
	// Declare a variable to hold a channel
	// Create the channel
	myChannel := make(chan string)
	go greeting(myChannel)
	receivedValue := <-myChannel
	fmt.Println(receivedValue)
	// fmt.Println(<-myChannel)
}
