// guess challenge players to guess a random number
package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand" // import path of package 'rand'
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	seconds := time.Now().Unix() // current time as an integer
	rand.Seed(seconds)

	target := rand.Intn(100) + 1
	fmt.Println("We've chosen a random number between 1 and 100")
	fmt.Println("Can you guess that the number")

	reader := bufio.NewReader(os.Stdin) // read keyboard input
	success := false

	for guesses := 0; guesses < 6; guesses++ {
		fmt.Println("You have", 6-guesses, "guesses left.")
		fmt.Println("Make a guess...")
		guess, err := reader.ReadString('\n') // read-up until press Enter
		if err != nil {
			log.Fatal(err)
		}
		guess = strings.TrimSpace(guess)
		number, err := strconv.Atoi(guess)
		if err != nil {
			log.Fatal(err)
		}

		if number < target {
			fmt.Println("Oops, Your guess was LOW.")
		} else if number > target {
			fmt.Println("Oops, Your guess was HIGH.")
		} else {
			success = true
			fmt.Println("Good job! You guessed it!")
			break
		}
	}

	if !success {
		fmt.Println("Sorry, you didn't guess the number. It was:", target)
	}
}
