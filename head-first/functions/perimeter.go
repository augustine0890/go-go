package main

import "fmt"

func calculatePerimeter(width float64, length float64) float64 {
	perimeter := (width + length) * 2
	return perimeter
}

func main() {
	total := calculatePerimeter(8.2, 10) + calculatePerimeter(5, 5.2) + calculatePerimeter(6.2, 4.5)
	fmt.Printf("You'll need %.2f meters of fencing\n", total)
}
