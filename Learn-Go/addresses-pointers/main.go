package main

import "fmt"

func brainwash(saying *string) {
	*saying = "Beep Boop"
}

// Pass the pointer of the variable
func changeValue(str *string) {
	*str = "changed!"
}

func changeValue2(str string) {
	str = "changed2!"
}

func main() {
	// Address
	treasure := "The friends we make along the way."
	var pointerForStr *string
	pointerForStr = &treasure
	fmt.Println(&treasure)
	fmt.Println(pointerForStr)

	// Pointers
	star := "Polaris"
	starAddress := &star
	fmt.Println("The address of star is", starAddress)

	num := 10
	fmt.Println("The address of num:", &num)

	*starAddress = "Sirius"
	fmt.Println("he actual value of star is", star)

	// Dereferencing
	lyrics := "Moments so dear"
	pointForLys := &lyrics
	*pointForLys = "Journeys to plan"
	fmt.Println(lyrics)
	// fmt.Println(*pointForLys)

	// Changing Values in Different Scopes
	greeting := "Hello there!"
	brainwash(&greeting)
	fmt.Println("greeting is now:", greeting)

	x := 7
	// fmt.Println(&x)
	y := &x
	fmt.Println(x, y)

	*y = 45
	fmt.Println(x, y)

	toChange := "hello"
	fmt.Println(toChange)
	changeValue(&toChange)
	fmt.Println(toChange)

	changeValue2(toChange)
	fmt.Println(toChange)

}
