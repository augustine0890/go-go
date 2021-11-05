package main

import "fmt"

type person struct {
	age int
}

func (p person) NewPerson(initialAge int) person {
	if initialAge < 0 {
		p.age = 0
		fmt.Println("Age must be greater than zero")
	}

	p.age = initialAge
	return p
}

func (p person) amIOld() {
	switch {
	case p.age < 13:
		fmt.Println("You're young.")
	case p.age > 13 && p.age < 18:
		fmt.Println("You're a teenager.")
	default:
		fmt.Println("You're old...")
	}
}

func (p person) yearPasses() person {
	p.age++
	return p
}

func main() {
	var T, age int

	fmt.Scan(&T)

	for i := 0; i < T; i++ {
		fmt.Scan(&age)
		p := person{age: age}
		p = p.NewPerson(age)
		p.amIOld()
		for j := 0; j < 3; j++ {
			p = p.yearPasses()
		}
		p.amIOld()
		fmt.Println()
	}
}
