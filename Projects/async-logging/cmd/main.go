package main

import (
	"flag"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	out := flag.String("out", "stdout", "File name to use for log output. If stdout is provided, then output is written directly to the console.")
	async := flag.Bool("async", false, "This flag determines if the logger should write asynchronously.")

	var w io.Writer
	var err error
	if strings.ToLower(*out) == "stdout" {
		w = os.Stdout
	} else {
		w, err := os.Create(*out)
		if err != nil {
			log.Fatal("Unable to open log file", err)
		}
	}
}
