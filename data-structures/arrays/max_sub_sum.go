package main

import (
	"fmt"
	"math"
)

// Calculate the maximum sum of n consecutive elements
func maxSubSum(arr []int, n int) int {
	if n > len(arr) {
		return math.MinInt8
	}
	maxSum, tempSum := 0, 0
	for i := 0; i < n; i++ {
		maxSum += arr[i]
	}
	tempSum = maxSum

	for i := n; i < len(arr); i++ {
		tempSum = tempSum + arr[i] - arr[i-n]
		maxSum = max(tempSum, maxSum)
	}
	return maxSum
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func main() {
	fmt.Println(maxSubSum([]int{2, 6, 9, 2, 1, 8, 5, 6, 3}, 3))

}
