package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	var now time.Time = time.Now()
	fmt.Println(now)
	var year int = now.Year()
	fmt.Println(year)

	broken := "G# r#ck!"
	replacer := strings.NewReplacer("#", "o")
	fixed := replacer.Replace(broken)
	fmt.Println(fixed)
}
