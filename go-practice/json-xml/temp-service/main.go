package main

import (
	"fmt"
	"temp-service/data"
	"temp-service/models"
	"temp-service/printer"
)

/**
- Read city data input from a JSON file
- Use a data structure to save cities
- Print out cities using the `range` operator
**/

func main() {
	fmt.Printf("Welcome to the Temperature Service!\n\n")

	cities, err := models.NewCities(data.NewReader())
	if err != nil {
		fmt.Println("Fatal error occurred: ", err)
		return
	}

	p := printer.New()
	defer p.Cleanup()
	p.CityHeader()

	for _, c := range cities.ListAll() {
		p.CityDetails(c)
	}
}
