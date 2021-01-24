package main

import "fmt"

// func main() {
// fmt.Println("Hello, world!")
// }

// P struct
type P struct {
	x string
	y int
}

func main() {
	b := P{"x", -1}
	a := [...]P{P{"a", 10},
		P{"b", 2},
		P{"c", 3}}
	for _, z := range a {
		if z.y > b.y {
			b = z
		}
	}
	fmt.Println(b.x)

	x := [...]int{1, 2, 3, 4, 5}
	y := x[0:2]
	z := x[1:4]
	fmt.Print(len(y), cap(y), len(z), cap(z))
}

// func main() {
// var xtemp int
// x1 := 0
// x2 := 1
// for x := 0; x < 5; x++ {
// xtemp = x2
// x2 = x2 + x1
// x1 = xtemp
// // fmt.Println(x2)
// }
// fmt.Println(x2)
// }
