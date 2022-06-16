package main

import "fmt"

// Munual Rotation: O(n)
func rightRotateM(arr []int, k int) []int {
	if len(arr) == 0 {
		k = 0
	} else {
		k = k % len(arr)
	}

	rotated := []int{}
	for i := len(arr) - k; i < len(arr); i++ {
		rotated = append(rotated, arr[i])
	}
	for i := 0; i < len(arr)-k; i++ {
		rotated = append(rotated, arr[i])
	}
	return rotated
}

// k = k % len(nums) to handle cases where k >= len(arr)
func rightRotate(arr []int, k int) {
	k = k % len(arr)
	reverse(arr, 0, len(arr)-1)
	reverse(arr, 0, k-1)
	reverse(arr, k, len(arr)-1)
}

// reverse an array in-place from startIndex to endIndex, inclusive
// O(n) time, O(1) space
func reverse(arr []int, start, end int) {
	for start <= end {
		arr[start], arr[end] = arr[end], arr[start]
		start++
		end--
	}
}

func main() {
	fmt.Println(rightRotateM([]int{10, 20, 30, 40, 50}, 3))

	arr := []int{10, 20, 30, 40, 50}
	rightRotate(arr, 3)
	fmt.Printf("Result: %d\n", arr)
}
