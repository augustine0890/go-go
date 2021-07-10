package main

import (
	"fmt"
	"log"
)

func calculatePerimeter(width float64, length float64) (float64, error) {
	if width < 0 {
		return 0, fmt.Errorf("a width of %.2f is invalid", width)
	}

	if length < 0 {
		return 0, fmt.Errorf("a length of %.2f is invalid", length)
	}
	perimeter := (width + length) * 2
	return perimeter, nil
}

func main() {
	p, err := calculatePerimeter(7.3, -12)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("You'll need %.2f meters of fencing\n", p)
}
