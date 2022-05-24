package main

import "fmt"

// Finds the smalles number in given array
// O(n)
func findMin(arr []int) int {
	min := arr[0]
	for _, v := range arr {
		if v < min {
			min = v
		}
	}
	return min
}

func main() {
	fmt.Println(findMin([]int{9, 2, 3, 6}))
	fmt.Println(findMin([]int{4, 2, 1, 5, 0}))
}
