package main

import "fmt"

func main() {
	fmt.Printf("%12s | %s\n", "Product", "Cost in Cents")
	fmt.Println("--------------------------------")

	fmt.Printf("%12s | %2d\n", "Stamps", 50)
	fmt.Printf("%12s | %2d\n", "Paper Clips", 5)
	fmt.Printf("%12s | %2d\n", "Tape", 99)

	fmt.Printf("%%7.3f: %7.3f\n", 12.3456)
	fmt.Printf("%%7.2f: %7.2f\n", 12.3456)
	fmt.Printf("%%.3f: %.3f\n", 12.3456)
	fmt.Printf("%%.2f: %.2f\n", 12.3456)
}
