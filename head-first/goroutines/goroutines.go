package main

import (
	"fmt"
	"time"
)

func a() {
	for i := 0; i < 50; i++ {
		fmt.Print("a")
	}
}

func b() {
	for i := 0; i < 50; i++ {
		fmt.Print("b")
	}
}

func repeat(s string) {
	for i := 0; i < 25; i++ {
		fmt.Print(s)
	}
}

func main() {
	go a()
	go b()

	go repeat("x")
	go repeat("y")

	time.Sleep(time.Second)
	fmt.Println("end main()")
}
