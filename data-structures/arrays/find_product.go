package main

import "fmt"

// Using nested loop
// O(n^2)
func findProduct(arr []int) []int {
	newArr := []int{}
	for i := 0; i < len(arr); i++ {
		temp := 1
		for j := 0; j < len(arr); j++ {
			if j != i {
				temp = temp * arr[j]
			}
		}
		newArr = append(newArr, temp)
	}
	return newArr
}

// Algorithm traverses over the array twice
// O(n)
func findProduct2(arr []int) []int {
	newArr := []int{}
	left, right := 1, 1

	for _, v := range arr {
		newArr = append(newArr, left)
		left = left * v
	}

	for i := len(arr) - 1; i >= 0; i-- {
		newArr[i] = newArr[i] * right
		right = right * arr[i]
	}

	return newArr
}

func main() {
	fmt.Println(findProduct([]int{1, 2, 3, 4}))
	fmt.Println(findProduct([]int{4, 2, 1, 5, 0}))

	fmt.Println(findProduct2([]int{1, 2, 3, 4}))
	fmt.Println(findProduct2([]int{4, 2, 1, 5, 0}))
}
