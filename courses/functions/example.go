package example

import "fmt"

func foo(y *int) {
	*y = *y + 1
}

// Slices contain a pointer to the array
func main() {
	x := 2
	foo(&x)
	fmt.Println(x)
}
