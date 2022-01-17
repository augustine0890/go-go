package main

import (
	"datatypes/organization"
	"fmt"
)

func main() {
	p := organization.NewPerson("Augustine", "Nguyen", organization.NewEuropeanUnionIdentifier("123-45-6789", "Spain"))
	err := p.SetTwitterHandler("@august")
	fmt.Printf("%T\n", organization.TwitterHandler("test"))
	if err != nil {
		fmt.Printf("An error occurred setting twitter handler: %s\n", err.Error())
	}

	fmt.Println(p.TwitterHandler())
	fmt.Println(p.TwitterHandler().RedirectUrl())
	fmt.Println(p.ID())
	fmt.Println(p.Country())
	fmt.Println(p.FullName())
}
