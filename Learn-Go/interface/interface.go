package main

import (
	"fmt"
	"math"
)

// Interfaces are named collections of method signatures
type geometry interface {
	area() float64
	perim() float64
}

type circle struct {
	radius float64
}

type rect struct {
	width, height float64
}

func (r rect) area() float64 {
	return r.width * r.height
}

func (r rect) perim() float64 {
	return 2 * (r.width + r.height)
}

func (c *circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c *circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

func main() {
	c := circle{4.5}
	r := rect{width: 5, height: 7}

	shapes := []geometry{&c, r}

	for _, shape := range shapes {
		// fmt.Println(shape.area())
		measure(shape)
	}
}
