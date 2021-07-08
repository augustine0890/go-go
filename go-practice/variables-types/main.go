package main

import "fmt"

func main() {
	fmt.Println(2235 * 1231)

	// Constants
	const earthsGravity = 9.80665
	fmt.Println(earthsGravity)

	// Variables
	var numOfFlavors int
	numOfFlavors = 57
	fmt.Println(numOfFlavors)

	var flavorScale float32 = 5.8
	fmt.Println(flavorScale)

	// Strings
	var favoriteSnack string
	favoriteSnack = "O'Start"
	fmt.Println("My favorite snack is " + favoriteSnack)

	// Zero values
	var emptyInt int8
	var emptyFloat float32
	var emptyString string
	var isClose bool
	fmt.Println(emptyInt, emptyFloat, emptyString, isClose)

	//Inferring Types
	daysOnVacation := 4
	var hourInDay = 24
	fmt.Println("You have spent", daysOnVacation * hourInDay, "hours on vacation.")

	// Default Type
	var cupsOfCoffeeConsumed uint
	cupsOfCoffeeConsumed = 11
	fmt.Println(cupsOfCoffeeConsumed)

	// Updating
	var basketTotal float64
	carrotPrice := 0.75

	basketTotal = basketTotal + carrotPrice
	fmt.Println(basketTotal)

	spinachPrice := 1.50
	// basketTotal = spinachPrice + basketTotal
	basketTotal += spinachPrice
	fmt.Println(basketTotal)

	coolSneakers := 65.99
	niceNecklace := 45.50
	
	var taxCalculation float64
	taxCalculation += coolSneakers
	taxCalculation += niceNecklace
	taxCalculation *= .08875
	fmt.Println("Purchase of", coolSneakers + niceNecklace,
				"with 8.875% sales tax", taxCalculation,
				"equal to", coolSneakers + niceNecklace + taxCalculation)

	// Multiple Var Declaration
	var magicNum, powerLevel int32
	magicNum = 2048
	powerLevel = 9001
	fmt.Println("magicNum is:", magicNum, "powerLevel is:", powerLevel)
	
	amount, unit := 10, "doll hairs"
	fmt.Println(amount, unit, ", that's expensive...")
}
