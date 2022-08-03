package main

import "fmt"

// Given an array nums of size n, return the majority element.
func majorityElement(nums []int) int {
	var countNums map[int]int = make(map[int]int)

	for _, num := range nums {
		if _, ok := countNums[num]; ok {
			countNums[num]++
		} else {
			countNums[num] = 1
		}
	}

	var maxCount int
	var value int
	for c := range countNums {
		if countNums[c] > maxCount {
			maxCount = countNums[c]
			value = c
		}
	}

	return value
}

func main() {

	fmt.Println(majorityElement([]int{3, 2, 3}))
	fmt.Println(majorityElement([]int{2, 2, 1, 1, 1, 2, 2}))
}
