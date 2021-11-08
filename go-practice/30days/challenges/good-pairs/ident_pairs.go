package main

import "fmt"

func numIdenticalPairs(nums []int) int {
	var count int

	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] == nums[j] {
				count++
			}
		}
	}
	return count
}

func identicalPairs(nums []int) int {
	result := 0
	count := make(map[int]int)

	for _, num := range nums {
		result += count[num]
		count[num]++
	}
	fmt.Printf("Count: %v\n", count)
	return result
}

func main() {
	res1 := numIdenticalPairs([]int{1, 2, 3, 1, 1, 3})
	fmt.Printf("Result: %d\n", res1)
	res2 := numIdenticalPairs([]int{1, 1, 1, 1})
	fmt.Printf("Result: %d\n", res2)
	res3 := numIdenticalPairs([]int{1, 2, 3})
	fmt.Printf("Result: %d\n", res3)

	res4 := identicalPairs([]int{1, 2, 3, 1, 1, 3})
	fmt.Printf("Result: %d\n", res4)
	res5 := identicalPairs([]int{1, 1, 1, 1})
	fmt.Printf("Result: %d\n", res5)
	res6 := identicalPairs([]int{1, 2, 3})
	fmt.Printf("Result: %d\n", res6)
}
