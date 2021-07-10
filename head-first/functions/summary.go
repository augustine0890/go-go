package main

import "fmt"

// func produce a table with student names, and the average of those scores
func scoreSummary(name string, score1 float64, score2 float64, score3 float64) {
	average := (score1 + score2 + score3) / 3
	fmt.Printf("%10s | %8.2f | %8.2f | %8.2f | %8.2f\n", name, score1, score2, score3, average)
}

func main() {
	fmt.Printf("%10s | %8s | %8s | %8s | %8s\n", "Name", "Score 1", "Score 2", "Score 3", "Average")
	fmt.Println("--------------------------------------------------------")
	scoreSummary("Jermaine", 95.4, 82.3, 74.6)
	scoreSummary("Jessie", 79.3, 99.1, 82.5)
	scoreSummary("Lamar", 82.2, 95.4, 77.6)
}
