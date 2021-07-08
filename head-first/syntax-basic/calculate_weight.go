package main

import "fmt"

func main() {
	var count int = 20
	unitWeights := 0.4
	totalWeights := float64(count) * unitWeights
	fmt.Println(count, "cans weight", totalWeights, "kilograms.")

	pebbleWeight := 0.1
	rockWeight := 1.2
	boulderWeight := 502.4
	totalWeight := pebbleWeight + rockWeight + boulderWeight
	fmt.Println(totalWeight)
}
