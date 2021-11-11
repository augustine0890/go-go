package main

import "fmt"

func lengthOfLongestSubstrings(s string) int {
	charMap := make(map[byte]int)
	longest := 0

	for l, r := 0, 0; r < len(s); r++ {
		if idx, ok := charMap[s[r]]; ok {
			l = max(l, idx)
		}
		longest = max(longest, r-l+1)
		charMap[s[r]] = r + 1
	}
	return longest
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(lengthOfLongestSubstrings("abcabcbb"))
	fmt.Println(lengthOfLongestSubstrings("abba"))
}
