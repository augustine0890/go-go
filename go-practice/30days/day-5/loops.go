package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	for i := 1; i < 11; i++ {
		res := n * i
		fmt.Printf("%v x %v = %v\n", n, i, res)
	}
}
