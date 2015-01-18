package models

import (
	"time"

	"appengine"
	"appengine/datastore"
)

type Link struct {
	Title       string
	Url         string
	Description string `datastore:",noindex"` // Markdown
	Tags        []string
	Created     time.Time
	Modified    time.Time
}

func NewLink(title string, url string, desc string, tags []string) *Link {
	e := new(Link)

	// User supplied content
	e.Title = title
	e.Url = url
	e.Description = desc
	e.Tags = tags

	// Computer generated content
	e.Created = time.Now()
	e.Modified = time.Now()

	return e
}

func (e *Link) Save(c appengine.Context) error {
	k := datastore.NewKey(c, "Link", e.Url, 0, nil)
	k2, err := datastore.Put(c, k, e)
	if err == nil {
		c.Infof("Wrote %+v", e)
		c.Infof("Old key: %+v; New Key: %+v", k, k2)
	} else {
		c.Warningf("Error writing entry: %v", e)
	}
	return err
}
