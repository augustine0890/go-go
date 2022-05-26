package main

import "fmt"

// Finds the smalles number in given array
// O(n)
func findMin(arr []int) int {
	min := arr[0]
	for _, v := range arr {
		if v < min {
			min = v
		}
	}
	return min
}

func mergeSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}
	first := mergeSort(arr[:len(arr)/2])
	second := mergeSort(arr[len(arr)/2:])
	return merge(first, second)
}

func merge(first, second []int) []int {
	final := []int{}
	i, j := 0, 0
	for i < len(first) && j < len(second) {
		if first[i] < second[j] {
			final = append(final, first[i])
			i++
		} else {
			final = append(final, second[j])
			j++
		}
	}

	for ; i < len(first); i++ {
		final = append(final, first[i])
	}

	for ; j < len(second); j++ {
		final = append(final, second[j])
	}

	return final
}

// mergeSort is in O(nlogn) --> takes O(nlogn) time
func findMin2(arr []int) int {
	sorted := mergeSort(arr)
	return sorted[0]
}

func main() {
	fmt.Println(findMin([]int{9, 2, 3, 6}))
	fmt.Println(findMin([]int{4, 2, 1, 5, 0}))

	fmt.Println(findMin2([]int{9, 2, 3, 6}))
	fmt.Println(findMin2([]int{4, 2, 1, 5, 0}))

}
