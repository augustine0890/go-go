package main

import (
	"fmt"
	"sort"
)

func canMakeArithmeticProgression(arr []int) bool {
	sort.Ints(arr)
	for i := 1; i < len(arr)-1; i++ {
		if arr[i]*2 != arr[i-1]+arr[i+1] {
			return false
		}
	}
	return true
}

func main() {
	res1 := canMakeArithmeticProgression([]int{3, 5, 1})
	fmt.Printf("Result: %t\n", res1)
	res2 := canMakeArithmeticProgression([]int{1, 2, 4})
	fmt.Printf("Result: %t\n", res2)
}
