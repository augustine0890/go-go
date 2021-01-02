package main

import (
	"fmt"
	"math"
	"time"
)

func eatTacos() {
	fmt.Println("Yum!")
}

func startGame() {
	instructions := "Press enter to start..."
	fmt.Println(instructions)
}

func isItLateInNewYork() string {
	var lateMesssage string
	t := time.Now()
	tz, _ := time.LoadLocation("America/New_York")
	nyHour := t.In(tz).Hour()
	if nyHour < 5 {
		lateMesssage = "Goodness it is late"
	} else if nyHour < 16 {
		lateMesssage = "It's not late at all!"
	} else if nyHour < 19 {
		lateMesssage = "I guess it's getting kind of late"
	} else {
		lateMesssage = "It's late"
	}

	return lateMesssage
}

func computeMarsYears(earthYears int) int {
	earthDays := earthYears * 365
	marsYears := earthDays / 687

	return marsYears
}

func specialComputation(x float64) float64 {
	return math.Log2(math.Sqrt(math.Tan(x)))
}

// GPA Multiple Return Values
func GPA(midtermGrade float32, finalGrade float32) (string, float32) {
	averageGrade := (midtermGrade + finalGrade) / 2
	var gradeLetter string

	if averageGrade > 90 {
		gradeLetter = "A"
	} else if averageGrade > 80 {
		gradeLetter = "B"
	} else if averageGrade > 70 {
		gradeLetter = "C"
	} else if averageGrade > 60 {
		gradeLetter = "D"
	} else {
		gradeLetter = "F"
	}

	return gradeLetter, averageGrade
}

func getLikesAndShares(postID int) (int, int) {
	var likesForPost, sharesForPost int
	switch postID {
	case 1:
		likesForPost = 5
		sharesForPost = 7
	case 2:
		likesForPost = 3
		sharesForPost = 11
	case 3:
		likesForPost = 22
		sharesForPost = 1
	case 4:
		likesForPost = 7
		sharesForPost = 9
	}
	fmt.Println("Likes: ", likesForPost, "Shares: ", sharesForPost)
	return likesForPost, sharesForPost
}

func queryDatabase(query string) string {
	var result string
	connectDatabase()
	// Add deferred call to disconnectDatabase() here
	defer disconnectDatabase()
	if query == "SELECT * FROM coolTable;" {
		result = "NAME|DOB\nVincent Van Gogh|March 30, 1853"
	}
	fmt.Println(result)
	return result
}

func connectDatabase() {
	fmt.Println("Connecting to the database.")
}

func disconnectDatabase() {
	fmt.Println("Disconnecting from the database.")
}

func main() {
	eatTacos()

	startGame()

	var nyLate string
	nyLate = isItLateInNewYork()
	fmt.Println(nyLate)

	myAge := 31
	myMartianAge := computeMarsYears(myAge)
	fmt.Println(myMartianAge)

	var a, b, c, d float64
	a = .0214
	b = 1.02
	c = 0.312
	d = 4.001

	a = specialComputation(a)
	b = specialComputation(b)
	c = specialComputation(c)
	d = specialComputation(d)
	fmt.Println(a, b, c, d)

	var myMidterm, myFinal float32
	myMidterm = 89.4
	myFinal = 74.9
	var myAverage float32
	var myGrade string
	myGrade, myAverage = GPA(myMidterm, myFinal)
	fmt.Println(myAverage, myGrade)

	var likes, shares int
	likes, shares = getLikesAndShares(4)
	if likes > 5 {
		fmt.Println("Woohoo! We got some likes.")
	}
	if shares > 10 {
		fmt.Println("We went viral!")
	}

	// Deferring Resolution
	queryDatabase("SELECT * FROM coolTable;")
}
