package main

import (
	"fmt"
	"strings"
)

/**
Make map to return at end
Loop over string, for each character
	- if the char is a number/letter and is a key in map, add one count
	- if not in map, set value to 1
	- if char is something other than --> do nothing
Return map
**/
func charCount(s string) map[string]int {
	counters := make(map[string]int)

	for _, c := range s {
		if isLetterNumeric(c) {
			char := strings.ToLower(string(c))
			if _, ok := counters[char]; ok {
				counters[char] += 1
			} else {
				counters[char] = 1
			}
		}
	}
	return counters
}

func isLetterNumeric(c rune) bool {
	return ('a' <= c && c <= 'z') || ('A' <= c && c <= 'Z') || (48 <= c && c <= 57)
}

func main() {
	fmt.Println(charCount("Your PIN number is: 0112329"))
	fmt.Println(charCount("hello"))
	fmt.Println(charCount("something you can count"))
}
