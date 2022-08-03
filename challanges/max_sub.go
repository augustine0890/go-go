package main

import (
	"fmt"
	"math"
)

// Given an integer array nums, find the contiguous subarray (containing at least one number) which has the largest sum and return its sum.
func maxSubArr(nums []int) int {
	maxSum := math.MinInt32
	currSum := 0

	for i := 0; i < len(nums); i++ {
		currSum += nums[i]
		if currSum > maxSum {
			maxSum = currSum
		}

		if currSum < 0 {
			currSum = 0
		}
	}

	return maxSum
}

func main() {
	fmt.Println(maxSubArr([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
	fmt.Println(maxSubArr([]int{-2, 1, -3, -4, -1, -2, 1, -5, 4}))
}
