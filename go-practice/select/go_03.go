package main

import "fmt"

func main() {
	ch := make(chan string)
	select {
	case v := <-ch:
		fmt.Println("received value", v)
	default:
		fmt.Println("default case executed")
	}
}
