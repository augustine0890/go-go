package main

import "fmt"

func quickSort(arr []int, low, high int) []int {
	if low < high {
		var p int
		arr, p = partition(arr, low, high)
		arr = quickSort(arr, low, p-1)
		arr = quickSort(arr, p+1, high)
	}
	return arr
}

func partition(arr []int, low, high int) ([]int, int) {
	pivot := arr[high]
	i := low
	for j := low; j < high; j++ {
		if arr[j] < pivot {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}
	arr[i], arr[high] = arr[high], arr[i]
	return arr, i
}

/**
Here are the steps:
1. Pick a pivot
2. Partition the array into two sub-arrays: ele less than pivot and ele greater than pivot
3. Call quicksort recursively on two sub-arrays
**/
func quickSortDC(arr []int) []int {
	if len(arr) < 2 {
		return arr
	} else {
		pivot := arr[0]
		var less, greater = []int{}, []int{}
		for _, num := range arr[1:] {
			if num < pivot {
				less = append(less, num)
			} else {
				greater = append(greater, num)
			}
		}
		less = append(quickSortDC(less), pivot)
		greater = quickSortDC(greater)
		return append(less, greater...)
	}
}

// The average runtime of quicksort is O(nlogn)
func main() {
	arr := []int{5, 6, 7, 2, 1, 0}
	fmt.Println(quickSort(arr, 0, len(arr)-1))
	fmt.Println(quickSortDC(arr))
}
