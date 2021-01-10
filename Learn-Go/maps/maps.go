package main

import "fmt"

func main() {
	var mp map[string]int = map[string]int{
		"apple":  5,
		"pear":   7,
		"orange": 8,
	}

	mp["apple"] = 9
	mp["augustine"] = 434
	fmt.Println(mp["apple"])
	fmt.Println(mp)
	fmt.Println(mp["apple"])

	delete(mp, "augustine")
	fmt.Println(len(mp))

	val, ok := mp["augustine"]
	fmt.Println(val, ok)

	// mp2 := make(map[string]int)
	// fmt.Println(mp2)
}
