package main

import "fmt"

func sumZero(arr []int) (int, int) {
	i, j := 0, len(arr)-1
	for i < j {
		sum := arr[i] + arr[j]
		if sum == 0 {
			return i, j
		} else if sum < 0 {
			i++
		} else {
			j--
		}
	}

	return -1, -1
}

func main() {
	fmt.Println(sumZero([]int{-4, -3, -2, -1, 0, 1, 2, 5}))
	fmt.Println(sumZero([]int{-4, -3, -2, -1, 0, 4}))
	fmt.Println(sumZero([]int{-4, -3, -2, -1, 0, 5}))
}
