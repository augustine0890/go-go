package main

import "fmt"

type location struct {
	x int
	y int
}

func isPathCrossing(path string) bool {
	x, y, tracker := 0, 0, make(map[location]bool)
	tracker[location{x, y}] = true

	for _, dir := range path {
		switch dir {
		case 'N':
			y++
		case 'S':
			y--
		case 'E':
			x++
		case 'W':
			x--
		}
		if tracker[location{x, y}] {
			return true
		}
		tracker[location{x, y}] = true
	}
	return false
}

func main() {
	res1 := isPathCrossing("NES")
	fmt.Printf("Result: %t\n", res1)
	res2 := isPathCrossing("NESWW")
	fmt.Printf("Result: %t\n", res2)
}
