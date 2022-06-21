package main

import "fmt"

// Inputs: sorted array, counts the unique values in the array
func countUnique(arr []int) int {
	if len(arr) == 0 {
		return 0
	}

	i := 0
	for j := 1; j < len(arr); j++ {
		if arr[i] != arr[j] {
			i++
			arr[i] = arr[j]
		}
	}
	return i + 1
}

func main() {
	fmt.Println(countUnique([]int{1, 1, 1, 1, 2}))
	fmt.Println(countUnique([]int{}))
	fmt.Println(countUnique([]int{1, 2, 3, 4, 4, 4, 7, 7, 12, 12, 13}))
	fmt.Println(countUnique([]int{-2, -1, -1, 0, 1}))
}
