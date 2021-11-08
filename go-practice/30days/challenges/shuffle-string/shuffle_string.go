package main

import "fmt"

func restoreString(s string, indices []int) string {
	shuffled := make([]byte, len(s))
	for idx, n := range indices {
		shuffled[n] = s[idx]
	}
	return string(shuffled)
}

func main() {
	res1 := restoreString("codeleet", []int{4, 5, 6, 7, 0, 2, 1, 3})
	fmt.Printf("Result: %s\n", res1)
	res2 := restoreString("abc", []int{0, 1, 2})
	fmt.Printf("Result: %s\n", res2)
	res3 := restoreString("aiohn", []int{3, 1, 4, 2, 0})
	fmt.Printf("Result: %s\n", res3)
	res4 := restoreString("aaiougrt", []int{4,0,2,6,7,3,1,5})
	fmt.Printf("Result: %s\n", res4)
	res5 := restoreString("art", []int{1,0,2})
	fmt.Printf("Result: %s\n", res5)
}
