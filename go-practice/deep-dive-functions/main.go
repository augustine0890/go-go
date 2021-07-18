package main

import (
	"deepdive/simplemath"
	"fmt"
)

func main() {
	fmt.Printf("%.2f\n", simplemath.Add(2, 5))
	answer, err := simplemath.Divide(6, 0)
	if err != nil {
		fmt.Printf("An error occurred: %s\n", err.Error())
	} else {
		fmt.Printf("%.2f\n", answer)
	}

	numbers := []float64{23.4, 34.2, 54.2}
	total := simplemath.Sum(numbers...)
	fmt.Printf("Total of sum: %.2f\n", total)
}
