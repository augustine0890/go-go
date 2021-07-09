package main

import "fmt"

func main() {
	for x := 1; x <= 3; x++ {
		fmt.Println("before continue")
		continue // skips back to top of loop
		// fmt.Println("after continue")
	}

	for x := 1; x <= 3; x++ {
		fmt.Println("before break")
		break // break out the loop
		// fmt.Println("after break")
	}
	fmt.Println("after loop")
}
