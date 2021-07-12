// Package: keyboard reads user input for the keyboardpackage
// Language: go
// Path: head-first/packages/keyboard/keyboard.go
package keyboard

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

// GetFloat reads a floating-point number from the keyboard.
// It returns the number read and any error encountered.
func getFloat() (float64, error) {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		return 0, err
	}

	input = strings.TrimSpace(input)             // remove whitespace (newlines, tabs, and regular spaces)
	number, err := strconv.ParseFloat(input, 64) // parse the string as a float64
	if err != nil {
		return 0, err
	}
	return number, nil
}
