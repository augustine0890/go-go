package main

import "fmt"

func mergeSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}
	first := mergeSort(arr[:len(arr)/2])
	second := mergeSort(arr[len(arr)/2:])

	return merge(first, second)
}

func merge(l []int, r []int) []int {
	final := []int{}
	i, j := 0, 0
	for i < len(l) && j < len(r) {
		if l[i] < r[j] {
			final = append(final, l[i])
			i++
		} else {
			final = append(final, r[j])
			j++
		}
	}

	for ; i < len(l); i++ {
		final = append(final, l[i])
	}
	for ; j < len(r); j++ {
		final = append(final, r[j])
	}
	return final
}

func binarySearch(arr []int, n int) int {
	var leftPointer = 0
	var rightPointer = len(arr) - 1
	for leftPointer <= rightPointer {
		var midPointer = int((leftPointer + rightPointer) / 2)
		var midValue = arr[midPointer]

		if midValue == n {
			return midPointer
		} else if midValue < n {
			leftPointer = midPointer + 1
		} else {
			rightPointer = midPointer - 1
		}
	}

	return -1
}

func main() {
	unsorted := []int{10, 6, 5, 1, 2, 8, 3, 4, 4, 7, 9}
	sorted := mergeSort(unsorted)
	fmt.Println(sorted)

	var arr = []int{3, 7, 23, 45, 57, 76}
	fmt.Println(binarySearch(arr, 45))
	fmt.Println(binarySearch(arr, 46))
}
