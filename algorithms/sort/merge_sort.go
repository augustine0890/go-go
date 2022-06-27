package main

import "fmt"

/**
Divide:
	- Divide the input slice into two (equal) halves
	- Recursively sort the two halves
Conquer:
	- Merge the two halves to form a sorted array
**/
func mergeSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}
	first := mergeSort(arr[:len(arr)/2])
	second := mergeSort(arr[len(arr)/2:])
	return merge(first, second)
}

// O(nlogn)
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

func main() {
	arr := []int{5, 6, 7, 2, 1, 0}
	fmt.Println(mergeSort(arr))

}
