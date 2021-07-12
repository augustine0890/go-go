package keyboard

import (
	"fmt"
	"log"
)

func ToCelsius() {
	fmt.Print("Enter the temperature in Fahrenheit: ")
	fahrenheit, err := getFloat()
	if err != nil {
		log.Fatal(err)
	}
	celsius := (fahrenheit - 32) * 5 / 9
	fmt.Printf("The temperature in Celsius is %.2f\n", celsius)
}
