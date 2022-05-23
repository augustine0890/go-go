package main

import "fmt"

func main() {
	sample1 := [2]int{1, 2}
	fmt.Printf("Sample 1 - Len: %d, %v\n", len(sample1), sample1)

	sample2 := [...]int{2, 3}
	fmt.Printf("Sample 2 - Len: %d, %v\n", len(sample2), sample2)

	sample3 := [2]int{}
	fmt.Printf("Sample 3 - Len: %d, %v\n", len(sample3), sample3)

	sample4 := [...]int{}
	fmt.Printf("Sample 4 - Len: %d, %v\n", len(sample4), sample4)

	letters := [3]string{"a", "b", "c"}
	length := len(letters)
	for i := 0; i < length; i++ {
		fmt.Println(letters[i])
	}

	for i, l := range letters {
		fmt.Printf("%d %s\n", i, l)
	}
}
