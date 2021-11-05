package main

import "fmt"

func main() {
	var T int

	fmt.Scan(&T)
	for i := 0; i < T; i++ {
		var s string
		fmt.Scan(&s)
		var even, odd string
		for idx, c := range s {
			if idx%2 == 0 {
				even += string(c)
			}
			if idx%2 != 0 {
				odd += string(c)
			}
		}
		fmt.Printf("%s %s\n", even, odd)
	}
}
