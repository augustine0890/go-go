package handlers

import (
	"fifa/data"
	"path"
	"path/filepath"
)

// reloads JSON into memory to ensure
// proper winner count during tests.
func setup() {
	p, _ := filepath.Abs("./../data/")
	fullpath := path.Join(p, "winners.json")
	data.LoadFromJSON(fullpath)
}
