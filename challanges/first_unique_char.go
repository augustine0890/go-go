package main

import "fmt"

// Given a string s, find the first non-repeating character in it and return its index. If it does not exist, return -1.

func firstUniqChar(s string) int {
	countStr := map[byte]int{}

	for i := 0; i < len(s); i++ {

		if _, ok := countStr[s[i]]; ok {
			countStr[s[i]] += 1
		} else {
			countStr[s[i]] = 1
		}
	}

	for j := 0; j < len(s); j++ {
		if countStr[s[j]] == 1 {
			return j
		}
	}

	return -1
}

func main() {
	fmt.Println(firstUniqChar("loveleetcode"))
}
