package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	// If statement
	heistReady := false // true
	if heistReady {
		fmt.Println("Let's Go!")
	} else {
		fmt.Println("Act normal.")
	}

	// Comparison Operators
	lockCombo := "2-35-19"
	robAttempt := "1-1-1"
	if lockCombo == robAttempt {
		fmt.Println("The vault is now opened.")
	}

	vaultAmt := 2358486
	if vaultAmt >= 200000 {
		fmt.Println("We're going to need more bags.")
	}

	// Logical Operators (And, Or)
	day := "Saturday"
	if day == "Saturday" || day == "Sunday" {
		fmt.Println("Enjoy the weekend!")
	} else {
		fmt.Println("Do some work.")
	}

	rightTime := true
	rightPlace := true
	if rightTime && rightPlace {
		fmt.Println("We're outta here!")
	} else {
		fmt.Println("Be patient...")
	}

	enoughRobbers := false
	enoughBags := true
	if enoughRobbers || enoughBags {
		fmt.Println("Grab everything!")
	} else {
		fmt.Println("Grab whatever you can!")
	}

	// Logical Operators (Not)
	readyToGo := true
	if !readyToGo {
		fmt.Println("Start the car!")
	} else {
		fmt.Println("What are we waiting for??")
	}

	// else if Statement
	amountStolen := 64650
	if amountStolen > 1000000 {
		fmt.Println("We've hit the jackpot!")
	} else if amountStolen >= 5000 {
		fmt.Println("Think of all the candy we can buy!")
	} else {
		fmt.Println("Why did we even do this?")
	}

	// The switch Statement
	clothingChoice := "sweater"
	switch clothingChoice {
	case "shirt":
		fmt.Println("We have shirts in S and M only.")
	case "polos":
		fmt.Println("We have polos in M, L, and XL.")
	case "sweater":
		fmt.Println("We have sweaters in S, M, L, and XL.")
	case "jackets":
		fmt.Println("We have jackets in all sizes.")
	default:
		fmt.Println("Sorry, we don't carry that")
	}

	// Scoped Short Declaration Statement
	if success := true; success {
		fmt.Println("We're rich!")
	} else {
		fmt.Println("Where did we go wrong?")
	}

	switch numOfThieves := 5; numOfThieves {
	case 1:
		fmt.Println("I'll take all $", amountStolen)
	case 2:
		fmt.Println("Everyone gets $", amountStolen/2)
	case 3:
		fmt.Println("Everyone gets $", amountStolen/3)
	case 4:
		fmt.Println("Everyone gets $", amountStolen/4)
	case 5:
		fmt.Println("Everyone gets $", amountStolen/5)
	default:
		fmt.Println("There's not enough to go around...")
	}

	// Randomizing
	rand.Seed(time.Now().UnixNano())
	amountLeft := rand.Intn(10000)
	fmt.Println("amountLeft if:", amountLeft)
	if amountLeft > 5000 {
		fmt.Println("What should I spend this on?")
  	} else {
    fmt.Println("Where did all my money go?")
  	}
}
