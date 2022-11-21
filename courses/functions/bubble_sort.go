package main

import "fmt"

func BubbleSort(nums []int) {
	for i := 0; i < len(nums)-1; i++ {
		for j := 0; j < len(nums)-1-i; j++ {
			Swap(nums, j)
		}
	}
}

func Swap(slices []int, index int) {
	if slices[index] > slices[index+1] {
		temp := slices[index+1]
		slices[index+1] = slices[index]
		slices[index] = temp
	}
}

func main() {
	nums := []int{15, 1, 4, 2, 8, 5, 11, 10, 3, 19, 16}
	BubbleSort(nums)
	fmt.Println(nums)
}
