package main

import "fmt"

func main() {
	fmt.Println("Enter number: ")
	var n int
	fmt.Scan(&n)
	switch {
	case n%2 != 0:
		fmt.Println("Weird")
	case 2 <= n, n <= 5, n%2 == 0:
		fmt.Println("Not Weird")
	case 6 <= n, n <= 20, n%2 == 0:
		fmt.Println("Weird")
	case n > 20, n%2 == 0:
		fmt.Println("Not Weird")
	default:
		fmt.Println("Not allowed.")
	}
}
