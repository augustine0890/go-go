package main

import (
	"fmt"
	"log"
	"math"
)

func squareRoot(number float64) (float64, error) {
	if number < 0 {
		return 0, fmt.Errorf("can't get square root of negative number.")
	}
	return math.Sqrt(number), nil
}

func main() {
	root1, err := squareRoot(9.0)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Printf("%0.2f\n", root1)
	}

	root2, err := squareRoot(-9.0)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Printf("%0.2f\n", root2)
	}

}
