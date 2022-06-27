package main

import "fmt"

/*
Given an array of integers sorted in increasing order and a target,
find the index of the first element in the array that is larger or equal to the target.
Assume that it is guaranteed to find a satisfying number.
- Time complexity: O(logn)
*/
func firstNotSmaller(arr []int, target int) int {
	left, right := 0, len(arr)-1
	index := -1
	for left <= right {
		mid := (left + right) / 2
		if arr[mid] >= target {
			index = mid
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return index
}

func main() {
	fmt.Println("Find first element :", firstNotSmaller([]int{1, 3, 3, 5, 8, 8, 10}, 2))
	fmt.Println("Find first element :", firstNotSmaller([]int{0}, 0))
	fmt.Println("Find first element :", firstNotSmaller([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 10))
	fmt.Println("Find first element :", firstNotSmaller([]int{1, 1, 1, 1, 4, 5}, 3))
}
