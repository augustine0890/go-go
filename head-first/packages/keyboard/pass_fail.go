package keyboard

import (
	"fmt"
	"log"
)

func FailPass() {
	fmt.Print("Enter a grade: ")
	grade, err := getFloat()
	if err != nil {
		log.Fatal(err)
	}
	var status string
	if grade >= 60 {
		status = "pass"
	} else {
		status = "fail"
	}
	fmt.Printf("The student's grade is %.2f and the student's status is %s.\n", grade, status)
}
