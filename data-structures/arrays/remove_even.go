package main

import "fmt"

// Remove even integers from slice.
func removeEven(arr []int) []int {
	odds := []int{}
	for i := 0; i < len(arr); i++ {
		if arr[i]%2 != 0 {
			odds = append(odds, arr[i])
		}
	}
	return odds
}

func main() {
	arr1 := []int{1, 2, 4, 5, 10, 6, 3}
	fmt.Println(removeEven(arr1))
	fmt.Println(removeEven([]int{3, 2, 41, 3, 34}))
}
