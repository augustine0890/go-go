package main

import "fmt"

/*
An array of boolean values is divided into two sections: the left section consists of all false, and the right section consists of all true.
Find the boundary of the right section, i.e. the index of the first true element.
If there is no true element, return -1.
- Time complexity: O(logn)
*/
func findBoundary(arr []bool) int {
	left, right := 0, len(arr)-1
	boundary := -1
	for left <= right {
		mid := (right + left) / 2
		if arr[mid] == false {
			left = mid + 1
		} else if arr[mid] == true {
			boundary = mid
			right = mid - 1
		}
	}
	return boundary
}

func main() {
	fmt.Println("Find boundary :", findBoundary([]bool{false, false, true, true, true}))
	fmt.Println("Find boundary :", findBoundary([]bool{true}))
	fmt.Println("Find boundary :", findBoundary([]bool{false, false, false}))
	fmt.Println("Find boundary :", findBoundary([]bool{false, false, true, true, true}))
	fmt.Println("Find boundary :", findBoundary([]bool{true, true, true, true, true}))
	fmt.Println("Find boundary :", findBoundary([]bool{false, true}))
}
