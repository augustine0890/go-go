package main

import "fmt"

func validAnagram(first, second string) bool {
	if len(first) != len(second) {
		return false
	}

	lookup := make(map[string]int)
	for _, s := range first {
		char := string(s)
		if _, ok := lookup[char]; ok {
			lookup[char] += 1
		} else {
			lookup[char] = 1
		}
	}

	seq2 := make(map[string]int)
	for _, s := range second {
		char := string(s)
		if _, ok := seq2[char]; ok {
			seq2[char] += 1
		} else {
			seq2[char] = 1
		}
	}

	for k, v := range lookup {
		char := string(k)
		if seq2[char] != v {
			return false
		}
	}

	return true
}

func main() {
	fmt.Println(validAnagram("", ""))                                     // true
	fmt.Println(validAnagram("aaz", "zza"))                               // false
	fmt.Println(validAnagram("anagram", "nagaram"))                       // true
	fmt.Println(validAnagram("rat", "car"))                               // false
	fmt.Println(validAnagram("awesome", "awesom"))                        // false
	fmt.Println(validAnagram("amanaplanacanalpanama", "acanalmanpamana")) // false
	fmt.Println(validAnagram("qwerty", "qeywrt"))                         // true
	fmt.Println(validAnagram("texttwisttime", "timetwisttext"))           // true
}
