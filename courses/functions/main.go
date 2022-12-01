package main

import "fmt"

func foo(y *int) {
	*y = *y + 1
}

func getMin(vals ...int) int {
	var minVal int = vals[0]
	for _, v := range vals {
		if v < minVal {
			minVal = v
		}
	}

	return minVal
}

// Slices contain a pointer to the array
func main() {
	defer func() {
		fmt.Println("Bye!")
	}()

	x := 2
	foo(&x)
	fmt.Println(x)

	nums := []int{15, 1, 4, 2, 8, 5, 11, 10, 3, 19, 16}
	BubbleSort(nums)
	fmt.Println(nums)

	mV := getMin(nums...)
	fmt.Println(mV)

	fmt.Println("Hello")
}
