package main

import (
	"fmt"
	"strings"
)

func main() {
	// strings.Builder value
	var builder strings.Builder

	builder.WriteRune('a') // adds rune onto end of string
	builder.WriteRune('b')
	builder.WriteRune('c')

	builder.WriteString("defg") // adds string

	fmt.Println(builder.String()) // returns string
}
