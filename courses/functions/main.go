package main

import "fmt"

func foo(y *int) {
	*y = *y + 1
}

// Slices contain a pointer to the array
func main() {
	x := 2
	foo(&x)
	fmt.Println(x)

	nums := []int{15, 1, 4, 2, 8, 5, 11, 10, 3, 19, 16}
	BubbleSort(nums)
	fmt.Println(nums)
}
