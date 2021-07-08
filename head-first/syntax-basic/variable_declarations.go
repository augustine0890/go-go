package main

import "fmt"

func main() {
	var quantity int
	var length, width float64
	var customerName string

	quantity = 4
	customerName = "Augustine Nguyen"
	length, width = 1.2, 2.4

	fmt.Println(customerName)
	fmt.Println("has ordered", quantity, "sheets")
	fmt.Println("each with an area of")
	fmt.Println(length*width, "square meters")
}
