package main

import (
	"appengine"
	"appengine/datastore"
	"fmt"
	"time"
)

type Entry struct {
	Id       uint64
	Title    string // optional
	Content  string // Markdown
	Datetime time.Time
	Created  time.Time
	Modified time.Time
	Tags     []string
	// TODO(icco): Define a meta field that is a json hash of extra data
}

func NewEntry(title string, content string, datetime time.Time, tags []string) *Entry {
	e := new(Entry)

	// User supplied content
	e.Title = title
	e.Content = content
	e.Datetime = datetime
	e.Tags = tags

	// Computer generated content
	e.Created = time.Now()
	e.Modified = time.Now()

	return e
}

func MaxId(c appengine.Context) (uint64, error) {
	entry := new(Entry)
	q := datastore.NewQuery("Entry").Order("-Id").Limit(1)
	_, err := q.Run(c).Next(&entry)
	if err != nil {
		return 0, nil
	} else {
		return entry.Id, nil
	}
}

func (e *Entry) hasId() bool {
	return (e.Id <= 0)
}

func (e *Entry) save(c appengine.Context) error {
	var k *datastore.Key
	if e.hasId() {
		id, err := MaxId(c)
		if err != nil {
			return err
		}
		e.Id = id + 1
		k = datastore.NewIncompleteKey(c, "Entry", nil)
	} else {
		k = datastore.NewKey(c, "Entry", fmt.Sprintf("%d", e.Id), 0, nil)
	}

	_, err := datastore.Put(c, k, e)
	return err
}
