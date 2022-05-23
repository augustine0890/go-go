package main

import (
	"fmt"
	"sort"
)

// Return two numbers add up to k
// O(n2)
func findSum(arr []int, k int) []int {
	for i := 0; i < len(arr)-1; i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i]+arr[j] == k {
				return []int{arr[i], arr[j]}
			}
		}
	}
	return []int{}
}

func binarySearch(arr []int, k int) bool {
	low := 0
	high := len(arr) - 1

	for low <= high {
		median := (low + high) / 2

		if arr[median] < k {
			low = median + 1
		} else {
			high = median - 1
		}
	}

	if low == len(arr) || arr[low] != k {
		return false
	}

	return true
}

func findSum2(arr []int, k int) []int {
	sort.Ints(arr)
	for i := 0; i < len(arr); i++ {
		index := binarySearch(arr, k-arr[i])
		if index {
			return []int{arr[i], k - arr[i]}
		}
	}
	return []int{}
}

func main() {
	fmt.Println(findSum([]int{1, 21, 3, 14, 5, 60, 7, 6}, 81))
	fmt.Println(findSum2([]int{1, 21, 3, 14, 5, 60, 7, 6}, 81))
	fmt.Println(findSum2([]int{1, 5, 3}, 2))
	fmt.Println(findSum2([]int{1, 2, 3, 4}, 5))

}
