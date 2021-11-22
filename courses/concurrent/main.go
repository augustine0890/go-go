package main

import (
	"fmt"
	"math/rand"
	"time"
)

var cache = map[int]Book{}
var rnd = rand.New(rand.NewSource(time.Now().UnixNano()))

func main() {
	for i := 0; i < 10; i++ {
		id := rnd.Intn(10) + 1
		go func(id int) {
			if b, ok := queryCache(id); ok {
				fmt.Println("from cache")
				fmt.Println(b)
			}
		}(id)

		go func(id int) {
			if b, ok := queryDatabase(id); ok {
				fmt.Println("from database")
				cache[id] = b
				fmt.Println(b)
			}
		}(id)
		// This is waiting for goroutine to finish (complete). The program execution completes when the main function exits
		// But we don't guarantee all of go routines to complete
		time.Sleep(time.Microsecond * 150)
		// fmt.Printf("Book not found with id: %v\n", id)
	}
}

func queryCache(id int) (Book, bool) {
	b, ok := cache[id]
	return b, ok
}

func queryDatabase(id int) (Book, bool) {
	time.Sleep(time.Microsecond * 300)
	for _, b := range books {
		if b.ID == id {
			return b, true
		}
	}

	return Book{}, false
}
