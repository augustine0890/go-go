package main

import (
	"arrays/datafile"
	"fmt"
	"log"
	"reflect"
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

func weekday() {
	weekday := [5]string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday"}
	for index, day := range weekday {
		fmt.Printf("%v %v\n", index, day)
	}
}

func counting(numbers interface{}) {
	arr := reflect.ValueOf(numbers)

	// We'll count the number of times each number occurs
	// within this array.
	counts := make(map[int64]int)

	for i := 0; i < arr.Len(); i++ {
		val := arr.Index(i).Int()
		counts[val]++
	}
	for number, count := range counts {
		fmt.Printf("%v occurred %v\n", number, count)
	}

}

func main() {
	fmt.Println("Go arrays")
	declare()

	fmt.Println("Go arrays puzzle")
	puzzle()

	fmt.Println("Go arrays weekday")
	weekday()

	fmt.Println("Go arrays counting")
	numberCount := []int64{1, 0, 2, 0, 1, 0, 0, 1, 2}
	counting(numberCount)

	numbersCount := []int64{4, 2, 6, 7, 8, 0, 1, 8, 7, 8,
		1, 5, 3, 2, 2, 1, 9, 6, 1, 0, 0, 0, 8, 4, 6,
		2, 2, 4, 7, 9, 6, 5, 9, 0, 5, 1, 1, 5, 4, 7,
		7, 9, 7, 8, 6, 3, 3, 3, 4, 8, 0, 4, 1, 1, 7,
		9, 8, 8, 1, 2, 3, 6, 4, 9, 2, 5, 8, 6, 7, 7,
		5, 4, 2, 9, 4, 4, 2, 2, 5, 5, 0, 0, 0, 9, 1,
		9, 5, 8, 0, 1, 1, 0, 5, 3, 8, 6, 3, 4, 4, 9}
	counting(numbersCount)

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
