package main

import "fmt"

// O(n^2)
// S(1)
func selectionSort(arr []int) []int {
	size := len(arr)
	for i := 0; i < size; i++ {
		min_idx := i
		for j := i + 1; j < size; j++ {
			if arr[min_idx] > arr[j] {
				min_idx = j
			}
			arr[i], arr[min_idx] = arr[min_idx], arr[i]
		}
	}
	return arr
}
func main() {
	fmt.Println(selectionSort([]int{12, 11, 13, 5, 6}))
	fmt.Println(selectionSort([]int{12, 5, 5, 10, 8, 1}))
	fmt.Println(selectionSort([]int{3, 5, 12, 8, 8, 1}))
	fmt.Println(selectionSort([]int{3, 5, 8, 10, 12, 1, 1}))

}
