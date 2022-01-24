package main

import (
	"datatypes/organization"
	"fmt"
	"reflect"
)

func main() {
	p := organization.NewPerson("Augustine", "Nguyen", organization.NewEuropeanUnionIdentifier("123-45-6789", "Spain"))
	err := p.SetTwitterHandler("@august")
	fmt.Printf("%T\n", organization.TwitterHandler("test"))
	if err != nil {
		fmt.Printf("An error occurred setting twitter handler: %s\n", err.Error())
	}

	// fmt.Println(p.TwitterHandler())
	// fmt.Println(p.TwitterHandler().RedirectUrl())
	// fmt.Println(p.ID())
	// fmt.Println(p.Country())
	// fmt.Println(p.FullName())

	eu1 := organization.NewEuropeanUnionIdentifier("12345", "France")
	eu2 := organization.NewEuropeanUnionIdentifier("12345", "France")

	if reflect.DeepEqual(eu1, eu2) {
		fmt.Println("We match")
	}

	map1 := map[string]int{
		"x": 1,
		"y": 2,
	}
	map2 := map[string]int{
		"x": 1,
		"y": 2,
		"z": 3,
	}
	fmt.Printf("Map equal: %v\n", reflect.DeepEqual(map1, map2))
}
