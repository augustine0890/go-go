package main

import (
	"fmt"
	"strings"
)

var monthMap = map[string]string{
	"Jan": "01",
	"Feb": "02",
	"Mar": "03",
	"Apr": "04",
	"May": "05",
	"Jun": "06",
	"Jul": "07",
	"Aug": "08",
	"Sep": "09",
	"Oct": "10",
	"Nov": "11",
	"Dec": "12",
}

func reformatDate(date string) string {
	dates := strings.Split(date, " ")
	year := dates[2]
	month := monthMap[dates[1]]
	day := dates[0][:2]
	if day[1] > byte('9') {
		day = "0" + dates[0][:1]
	}
	formated := fmt.Sprintf("%s-%s-%s", year, month, day)
	return formated
}

func main() {
	res1 := reformatDate("20th Oct 2052")
	fmt.Printf("Result: %s\n", res1)
	res2 := reformatDate("6th Jun 1933")
	fmt.Printf("Result: %s\n", res2)
	res3 := reformatDate("26th May 1960")
	fmt.Printf("Result: %s\n", res3)
}
