package main

import "fmt"

// O(n^2)
// S(1)
func inserSort(arr []int) []int {
	size := len(arr)

	for i := 1; i < size; i++ {
		key := arr[i]
		j := i - 1
		for j >= 0 {
			if arr[j] > key {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
			j--
		}
	}
	return arr
}

func main() {
	fmt.Println(inserSort([]int{12, 11, 13, 5, 6}))
	fmt.Println(inserSort([]int{12, 5, 5, 10, 8, 1}))
	fmt.Println(inserSort([]int{3, 5, 12, 8, 8, 1}))
	fmt.Println(inserSort([]int{3, 5, 8, 10, 12, 1, 1}))

}
