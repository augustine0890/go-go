package main

import "fmt"

func main() {
	var length float64 = 1.2
	var width int = 2

	fmt.Println("Area is", length*float64(width))
	fmt.Println("length > width", length > float64(width))

	var price int = 100
	fmt.Println("Price is", price, "dollars.")

	var taxRate float64 = 0.08
	var tax float64 = float64(price) * taxRate
	fmt.Println("Tax is", tax, "dollars.")

	var total float64 = float64(price) + tax
	fmt.Println("Total cost is", total, "dollars.")

	var availableFunds int = 120
	fmt.Println(availableFunds, "dollars available.")
	fmt.Println("Within budget?", total <= float64(availableFunds))
}
