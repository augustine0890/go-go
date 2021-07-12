package main

import (
	"fmt"
	"hello/calc"
	"hello/greeting"
	"hello/keyboard"
)

func main() {
	greeting.Hello()
	greeting.Hi()

	fmt.Println(calc.Add(1, 2))
	fmt.Println(calc.Subtract(7, 5))

	// keyboard
	fmt.Println("Reports whether a grade is passing or failing")
	keyboard.FailPass()

	fmt.Println("Converts a temperature to Celsius")
	keyboard.ToCelsius()
}
