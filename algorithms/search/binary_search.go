package main

import "fmt"

// Binary search only works when your list is in sorted order
func binarySearch(arr []int, k int) int {
	low, high := 0, len(arr)-1
	for low <= high {
		mid := low + (high-low)/2
		guess := arr[mid]
		if guess == k {
			return mid
		} else {
			if guess < k {
				low = mid + 1
			} else {
				high = mid - 1
			}
		}
	}
	return -1
}

func main() {
	fmt.Println(binarySearch([]int{1, 3, 5, 7, 9}, 3))
	fmt.Println(binarySearch([]int{1, 3, 5, 7, 9}, 4))
	fmt.Println(binarySearch([]int{1, 3, 5, 7, 9}, 9))
}
