package main

import (
	"fmt"
	"os"
)

/*
Defer is used to ensure that a function call is performed later in a
program's execution.
*/
func createFile(p string) *os.File {
	fmt.Println("creating")
	file, err := os.Create(p)
	if err != nil {
		panic(err)
	}
	return file
}

func writeFile(f *os.File) {
	fmt.Println("writing")
	fmt.Fprintln(f, "data")
}

func closeFile(f *os.File) {
	fmt.Println("closing")
	err := f.Close()

	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}

func main() {
	f := createFile("/tmp/defer.txt")
	defer closeFile(f)
	writeFile(f)
}
