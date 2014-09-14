package main

import (
	"appengine"
	"appengine/datastore"
	"fmt"
	"time"
)

type Entry struct {
	Id       int64
	Title    string // optional
	Content  []byte // Markdown
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
	e.Content = []byte(content)
	e.Datetime = datetime
	e.Tags = tags

	// Computer generated content
	e.Created = time.Now()
	e.Modified = time.Now()

	return e
}

func GetEntry(c appengine.Context, id int64) (*Entry, error) {
	var entry Entry
	q := datastore.NewQuery("Entry").Filter("Id =", id)
	_, err := q.Run(c).Next(&entry)
	if err != nil {
		c.Warningf("Error getting entry %d", id)
		return nil, err
	}
	return &entry, nil
}

func MaxId(c appengine.Context) (int64, error) {
	var entry Entry
	q := datastore.NewQuery("Entry").Order("-Id").Limit(1)
	_, err := q.Run(c).Next(&entry)
	if err != nil {
		return 0, err
	}
	return entry.Id, nil
}

func AllPosts(c appengine.Context) (*[]Entry, error) {
	q := datastore.NewQuery("Entry").Order("-Datetime")
	entries := new([]Entry)
	_, err := q.GetAll(c, entries)
	return entries, err
}

func (e *Entry) hasId() bool {
	return (e.Id <= 0)
}

func (e *Entry) save(c appengine.Context) error {
	var k *datastore.Key
	if e.hasId() {
		id, _ := MaxId(c)
		e.Id = id + 1
		k = datastore.NewIncompleteKey(c, "Entry", nil)
	} else {
		k = datastore.NewKey(c, "Entry", fmt.Sprintf("%d", e.Id), 0, nil)
	}

	_, err := datastore.Put(c, k, e)
	if err == nil {
		c.Infof("Wrote %+v", e)
	} else {
		c.Warningf("Error writing entry: %v", e)
	}
	return err
}
