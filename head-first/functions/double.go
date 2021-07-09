package main

import "fmt"

func double(number *int) { // take pointer type int
	*number *= 2 // value at the pointer
}

func main() {
	amount := 6
	double(&amount)      // pass the pointer instead of the value
	fmt.Println(amount)  // value
	fmt.Println(&amount) // address
}
