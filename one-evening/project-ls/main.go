package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	pwd, err := os.Getwd()
	if err != nil {
		os.Exit(1)
	}

	dir := pwd + "/testdata"
	fmt.Println(strings.Join(listFiles(dir), " "))
}

func listFiles(dirname string) []string {
	var dirs []string

	files, _ := ioutil.ReadDir(dirname)

	for _, f := range files {
		dirs = append(dirs, f.Name())
	}

	return dirs
}
