package main

import (
	"fmt"
	"hello/calc"
	"hello/greeting"
)

func main() {
	greeting.Hello()
	greeting.Hi()

	fmt.Println(calc.Add(1, 2))
	fmt.Println(calc.Subtract(7, 5))
}
