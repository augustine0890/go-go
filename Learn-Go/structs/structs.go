package main

import "fmt"

type person struct {
	name string
	age  int
}

func newPerson(name string) *person {
	p := person{name: name}
	p.age = 42
	return &p
}

func main() {
	fmt.Println(person{"Bob", 30})
	fmt.Println(person{name: "Augustine", age: 31})
	fmt.Println(person{name: "Fred"})

	fmt.Println(&person{name: "Ann", age: 40})
	fmt.Println(newPerson("John"))

	s := person{name: "Sean", age: 50}
	fmt.Println(s.name)

	// An & prefix yields a pointer to the struct.
	sp := &s
	fmt.Println(sp.age)

	// Use dots with struct pointers - the pointers are automatically dereferenced.
	sp.age = 51
	fmt.Println(sp.age)
	fmt.Println(s.age)

}
