package main

import "fmt"

// O(nlogn)
func findSecondMax(arr []int) int {
	if len(arr) >= 2 {
		sorted := mergeSort(arr)
		return sorted[len(arr)-2]
	}
	return -1
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

//  O(n) since the array is traversed twice.
func findSecondMax2(arr []int) int {
	var max int
	var secondMax int
	for _, v := range arr {
		if v > max {
			max = v
		}
	}

	for _, v := range arr {
		if v != max && v > secondMax {
			secondMax = v
		}
	}

	return secondMax
}

func main() {
	fmt.Println(findSecondMax([]int{9, 2, 3, 6}))
	fmt.Println(findSecondMax([]int{4, 2, 1, 5, 0}))

	fmt.Println(findSecondMax2([]int{9, 2, 3, 6}))
	fmt.Println(findSecondMax2([]int{4, 2, 1, 5, 0}))

}
