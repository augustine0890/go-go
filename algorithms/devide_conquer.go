package main

import "fmt"

func sum(arr []int) int {
	if len(arr) == 0 {
		return 0
	}
	return arr[0] + sum(arr[1:])
}

func count(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	return 1 + count(nums[1:])
}

func max(nums []int) int {
	if len(nums) == 2 {
		if nums[0] > nums[1] {
			return nums[0]
		}
		return nums[1]
	}
	subMax := max(nums[1:])
	if nums[0] > subMax {
		return nums[0]
	}
	return subMax

}

func main() {
	fmt.Println(sum([]int{1, 2, 3, 4}))
	fmt.Println(count([]int{1, 2, 3, 4}))
	fmt.Println(max([]int{1, 5, 10, 25, 16, 1}))
}
