package main

import "fmt"

func twoSum(nums []int, target int) []int {
	pair := make(map[int]int)

	for idx, n := range nums {
		sub := target - n
		_, found := pair[sub]
		if found {
			return []int{pair[sub], idx}
		}
		pair[nums[idx]] = idx
	}
	return nil
}

func main() {
	nums := []int{1, 3, 7, 9, 2}
	res := twoSum(nums, 11)
	fmt.Printf("%v\n", res)
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
	fmt.Println(twoSum([]int{3, 2, 4}, 6))
	fmt.Println(twoSum([]int{3, 2, 4}, 9))
}
