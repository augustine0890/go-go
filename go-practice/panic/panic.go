package main

import (
	"fmt"
	"log"
)

/*
A panic typically means something went unexpectedly wrong.
Mostly we use it to fail fast on errors that shouldn't occur during normal operation
Don't use when file can't be opened, unless it is critical
Use for unrecoverable events - cannot obtain TCP port for web server
Recover:
	- Used to recover from panics
	- Only useful in deferred functions
	- Current function will not attempt to continue, but higher functions in call stack will
*/
func panicker() {
	fmt.Println("about to panic")
	defer func() {
		if err := recover(); err != nil {
			log.Println("Error:", err)
			panic(err)
		}
	}()
	panic("something bad happened")
	// fmt.Println("done panicking")
}

func main() {
	fmt.Println("start")
	panicker()
	fmt.Println("end")
}
