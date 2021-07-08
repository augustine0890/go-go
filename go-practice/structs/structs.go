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

// Student methods defined on struct types
type Student struct {
	name   string
	grades []int
	age    int
}

func (s Student) getAge() int {
	return s.age
}

func (s *Student) setAge(age int) {
	s.age = age
}

func (s Student) getAverageGrade() float32 {
	sum := 0
	for _, v := range s.grades {
		sum += v
	}
	return float32(sum) / float32(len(s.grades))
}

func (s *Student) getMaxGrade() int {
	curMax := 0
	for _, v := range s.grades {
		if curMax < v {
			curMax = v
		}
	}
	return curMax
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

	s1 := Student{name: "Augustine", grades: []int{80, 78, 92, 95}, age: 21}
	s2 := Student{name: "John", grades: []int{80, 79, 78, 92, 99, 95}, age: 23}
	fmt.Println("Student:", s1)
	fmt.Println("Age Before:", s1.getAge())
	s1.setAge(22)
	fmt.Println("Age After:", s1.age)

	average1 := s1.getAverageGrade()
	fmt.Println("Average Grade:", average1)
	fmt.Println("Max Grade:", s1.getMaxGrade())
	average2 := s2.getAverageGrade()
	fmt.Println("Average Grade:", average2)
	fmt.Println("Max Grade:", s2.getMaxGrade())

}
