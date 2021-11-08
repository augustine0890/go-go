package main

import "fmt"

func countOdds(low int, high int) int {
	if low % 2 != 0 || high % 2 != 0 {
		return (high - low) / 2 + 1
	}
	return (high - low) / 2
}

func main() {
	res1 := countOdds(3, 7)
	fmt.Printf("Result: %d\n", res1)
	res2 := countOdds(8, 10)
	fmt.Printf("Result: %d\n", res2)
}
