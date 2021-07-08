package main

import "fmt"

func main() {
	var a []int = []int{1, 3, 5, 25, 52, 23, 87, 25}

	// for i := 0; i < len(a); i++ {
	// fmt.Println(a[i])
	// }

	// for i, el := range a {
	// fmt.Printf("%d: %d\n", i, el)
	// }

	// for _, el := range a {
	// fmt.Println(el)
	// }

	for i, el1 := range a {
		for j := i + 1; j < len(a); j++ {
			el2 := a[j]
			if el2 == el1 {
				fmt.Println(el1)
			}
		}
	}
}
