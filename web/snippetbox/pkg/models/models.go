package models

import (
	"errors"
	"time"
)

var ErrNoRecord = errors.New("models: no matching record found")

type Snippet struct {
	ID      int
	Title   string
	Content string
	Created time.Time
	Expires time.Time
}

// format the date
func (snippet *Snippet) FormatCreatedDate() string {
	return snippet.Created.Format(time.RFC1123)
}

func (snippet *Snippet) FormatExpiresDate() string {
	return snippet.Expires.Format(time.RFC1123)
}
