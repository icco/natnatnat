package main

import (
	"time"
)

type Entry struct {
	id       uint64
	title    string // optional
	content  string // Markdown
	datetime time.Time
	created  time.Time
	modified time.Time
	tags     []string
	// TODO(icco): Define a meta field that is a json hash of extra data
}

func WriteEntry(title string, content string, datetime time.Time, tags []string) {

}
