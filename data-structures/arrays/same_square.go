package main

import "fmt"

// O(n) time, O(n) space
func same(arr1, arr2 []int) bool {
	if len(arr1) != len(arr2) {
		return false
	}

	square := make(map[int]int)
	for _, v := range arr1 {
		square[findSquare(v)] = v
	}
	for _, v := range arr2 {
		if _, ok := square[v]; !ok {
			return false
		}
	}

	return true
}

func findSquare(x int) int {
	return x * x
}

func main() {
	fmt.Println(same([]int{1, 2, 3, 2}, []int{9, 1, 4, 4}))
	fmt.Println(same([]int{1, 2, 3, 2}, []int{9, 1, 4}))
	fmt.Println(same([]int{1, 2, 3, 2, 4}, []int{9, 1, 4, 3, 16}))
}
