package main

import (
	"fmt"
	"log"
	"reflect"
	"sort"
	"time"

	"anon/authenticator"
)

func printer() func() {
	k := 1
	return func() {
		fmt.Printf("Print n. %d\n", k)
		k++
	}
}

/*
Function that don't have a function name
*/
func main() {
	func() {
		fmt.Println("I'm an anonymous function")
	}()

	// func literal not executed
	myFunc := func() int {
		fmt.Println("func literal")
		return 32
	}
	fmt.Println(reflect.TypeOf(myFunc))

	// func literal invoked
	funcValue := func() int {
		fmt.Println("func literal invoked")
		return 32
	}()
	fmt.Println(reflect.TypeOf(funcValue))

	p := printer()
	p()
	p()
	p()

	// func params
	scores := []int{10, 79, 67, 4, 38, 24, 87}
	sort.Slice(scores, func(i, j int) bool { return scores[i] < scores[j] })
	fmt.Println(scores)

	sort.Slice(scores, func(i, j int) bool { return scores[i] > scores[j] })
	fmt.Println(scores)

	// Authenticator
	auth := authenticator.New("test", authenticator.UnverifiedEmailAllowed(), authenticator.WithCustomJwtDuration(time.Second))

	if auth.IsValidJWT("invalid") {
		log.Println("Hmm... try to use another library.")
	}
}
