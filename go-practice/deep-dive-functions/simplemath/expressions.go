package simplemath

import (
	"errors"
	"math"
)

func Add(num1, num2 float64) float64 {
	return num1 + num2
}

func Subtract(num1, num2 float64) float64 {
	return num1 - num2
}

func Multiply(num1, num2 float64) float64 {
	return num1 * num2
}

func Divide(num1, num2 float64) (answer float64, err error) {
	if num2 == 0 {
		return math.NaN(), errors.New("Cannot divide by zero.")
	}

	answer = num1 / num2
	return
}

func Sum(values ...float64) float64 {
	var total float64
	for _, value := range values {
		total += value
	}
	return total
}
