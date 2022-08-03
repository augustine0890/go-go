package main

import "fmt"

// Given an integer array nums, move all 0's to the end of it while maintaining the relative order of the non-zero elements.
func moveZeroes(nums []int) []int {
	c := 0

	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[c] = nums[i]
			c++
		}
	}

	for ; c < len(nums); c++ {
		nums[c] = 0
	}

	return nums

}

func main() {
	fmt.Println(moveZeroes([]int{0, 1, 0, 3, 12}))
}
