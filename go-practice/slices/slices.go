package main

import "fmt"

func main() {
	var x [5]int = [5]int{1, 2, 3, 4, 5}
	var s []int = x[1:3]
	fmt.Println(s)
	fmt.Println(len(s))
	fmt.Println(cap(s))
	fmt.Println(s[:cap(s)])

	var a []int = []int{5, 6, 7, 8, 9}
	fmt.Println(cap(a))
	fmt.Println(cap(a[:3]))
	b := append(a, 10)
	fmt.Println(b)
	d := append(b, []int{11, 12, 13, 14}...)
	fmt.Println(d)
	f := append(d[:5], d[6:]...)
	fmt.Println(f)

	c := make([]int, 5)
	fmt.Println(c)
	fmt.Printf("%T\n", c)
}
