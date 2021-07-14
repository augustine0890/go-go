package main

import (
	"arrays/datafile"
	"fmt"
	"log"
)

func declare() {
	var numbers [3]int
	numbers[0] = 43
	numbers[1] = 21
	numbers[2] = 17
	fmt.Printf("%#v\n", numbers)

	var letters = [3]string{"a", "b", "c"}
	fmt.Printf("%#v\n", letters)

	for i := 0; i < len(letters); i++ {
		fmt.Printf("index %v: %v\n", i, letters[i])
	}

	var notes = [7]string{"do", "re", "mi", "fa", "so", "la", "si"}
	for index, note := range notes {
		fmt.Printf("index %v: %v\n", index, note)
	}
}

func puzzle() {
	numbers := [6]int{3, 16, -2, 10, 24, 14}
	for index, number := range numbers {
		if number >= 10 && number <= 20 {
			fmt.Println(index, number)
		}
	}
}

func main() {
	fmt.Println("Go arrays")
	declare()

	fmt.Println("Go arrays puzzle")
	puzzle()

	numbers, err := datafile.GetFloats("data.txt")
	if err != nil {
		log.Fatal(err)
	}
	var sum float64 = 0.0
	for _, num := range numbers {
		sum += num
	}
	sampleCount := float64(len(numbers))
	fmt.Printf("Average: %.2f\n", sum/sampleCount)
}
