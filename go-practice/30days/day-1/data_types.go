package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var i int64 = 4
	var d float64 = 4.0
	var s string = "30 days "

	scanner := bufio.NewScanner(os.Stdin)

	// Declare second integer, double, ans String variable.
	var (
		ii int64
		id float64
		is string
	)

	// Read and save an integer, double, ans String variables.
	fmt.Scan(&ii)
	fmt.Scan(&id)
	scanner.Scan()
	is = scanner.Text()
	fmt.Println(i + ii)
	fmt.Printf("%.1f\n", d+id)
	fmt.Println(s + is)
}
