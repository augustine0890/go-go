package main

import "fmt"

func numWaterBottles(numBottles int, numExchanges int) int {
	var total int
	var emptyBottles int
	total = numBottles
	emptyBottles = numBottles

	for emptyBottles >= numExchanges {
		exchanged := emptyBottles / numExchanges
		emptyBottles = exchanged + emptyBottles % numExchanges

		total += exchanged
	}

	return total
}

func main() {
	res1 := numWaterBottles(9, 3)
	fmt.Printf("Result: %d\n", res1)
	res2 := numWaterBottles(15, 4)
	fmt.Printf("Result: %d\n", res2)
	res3 := numWaterBottles(5, 5)
	fmt.Printf("Result: %d\n", res3)
	res4 := numWaterBottles(2, 3)
	fmt.Printf("Result: %d\n", res4)
}
