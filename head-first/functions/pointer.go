package main

import (
	"fmt"
	"reflect"
)

/*
The type of a pointer is written with a * symbol.
The type of a pointer to an `int` variable is `*int`
`*myIntPointer` is value at myIntPointer
*/

func createPointer() *float64 {
	var myFloat = 98.5
	return &myFloat // pointer of type float64
}

func printPointer(myBoolPointer *bool) { // pointer type
	fmt.Println(*myBoolPointer)
}

/*
Get a pointer to a variable by `address of` (&) right before the variable name
Access the value of pointer by putting a `*` right before the pointer name `*myPointer`
*/
func main() {
	var myInt int
	fmt.Println(&myInt)
	fmt.Println(reflect.TypeOf(&myInt)) // get pointer of myInt
	var myIntPtr *int
	myIntPtr = &myInt
	fmt.Println(myIntPtr)

	var myFloat float64
	fmt.Println(reflect.TypeOf(&myFloat)) // get pointer of myFloat
	myFloatPtr := &myFloat
	fmt.Println(myFloatPtr)

	var myBool bool
	fmt.Println(reflect.TypeOf(&myBool))

	myString := "hello"
	myStringPtr := &myString
	fmt.Println(myStringPtr)
	fmt.Println(*myStringPtr) // value at myStringPtr --> "hello"
	fmt.Println(myString)     // "hello"
	*myStringPtr = "hi"
	fmt.Println(*myStringPtr) // value at myStringPtr after reassign --> "hi"
	fmt.Println(myString)     // "hi"

	var myFloatPointer *float64 = createPointer()
	fmt.Println(*myFloatPointer) // value at myFloatPointer

	var valBool = true
	printPointer(&valBool) // pass pointer to function
}
