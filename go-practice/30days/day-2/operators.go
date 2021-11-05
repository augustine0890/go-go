package main

import (
	"fmt"
	"math"
)

func main() {
	var (
		mealCost   float64
		tipPercent float64
		taxPercent float64
	)

	fmt.Scan(&mealCost, &tipPercent, &taxPercent)
	tip := mealCost * tipPercent / 100
	tax := mealCost * taxPercent / 100
	totalCost := mealCost + tip + tax
	fmt.Printf("The total meal cost is %v dollars.\n", math.Round(totalCost))
}
