package main

import "fmt"

// O(n) time
func rearrange(arr []int) []int {
	var neg, pos, re []int
	for _, e := range arr {
		if e < 0 {
			neg = append(neg, e)
		} else {
			pos = append(pos, e)
		}
	}
	re = append(neg, pos...)
	return re
}

// O(n) time, O(1) space
func rearrangeInPlace(arr []int) []int {
	leftMostPos := 0
	for idx := 0; idx < len(arr); idx++ {
		if arr[idx] < 0 {
			if idx != leftMostPos {
				// swap the two
				arr[idx], arr[leftMostPos] = arr[leftMostPos], arr[idx]
			}
			leftMostPos += 1
		}
	}
	return arr
}

func rightArrange(arr []int) []int {
	rightMostNeg := len(arr) - 1
	for idx := len(arr) - 1; idx >= 0; idx-- {
		if arr[idx] > 0 {
			if idx != rightMostNeg {
				arr[idx], arr[rightMostNeg] = arr[rightMostNeg], arr[idx]
			}
			rightMostNeg -= 1
		}
	}
	return arr
}

func main() {
	fmt.Println(rearrange([]int{10, -1, 20, 4, 5, -9, -6}))
	fmt.Println(rearrange([]int{7, -1, 11, 4, -8, -19, -16}))

	fmt.Println(rearrangeInPlace([]int{10, -1, 20, 4, 5, -9, -6}))
	fmt.Println(rearrangeInPlace([]int{7, -1, 11, 4, -8, -19, -16}))

	fmt.Println(rightArrange([]int{10, -1, 20, 4, 5, -9, -6}))
	fmt.Println(rightArrange([]int{7, -1, 11, 4, -8, -19, -16}))

}
