package main

import "fmt"

// Merges two sorted arrays of m and n elements --> into another sorted array.
func mergeArrays(arr1, arr2 []int) []int {
	result := []int{}
	i, j := 0, 0

	for i < len(arr1) && j < len(arr2) {
		if arr1[i] < arr2[j] {
			result = append(result, arr1[i])
			i++
		} else {
			result = append(result, arr2[j])
			j++
		}
	}

	for ; i < len(arr1); i++ {
		result = append(result, arr1[i])
	}
	for ; j < len(arr2); j++ {
		result = append(result, arr2[j])
	}

	return result
}

// Time complexity is O(n+m) wher n and m are length of the arrays.
func main() {
	fmt.Println(mergeArrays([]int{1, 3, 4, 5}, []int{2, 6, 7, 8}))
	fmt.Println(mergeArrays([]int{4, 5, 6}, []int{-2, -1, 0, 7}))
}
