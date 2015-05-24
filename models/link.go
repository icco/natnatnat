package models

import (
	"fmt"
	"strings"
	"time"

	"golang.org/x/net/context"

	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
	"google.golang.org/appengine/log"
)

type Link struct {
	Title       string
	Url         string
	Description string `datastore:",noindex"` // Markdown
	Tags        []string
	Posted      time.Time
	Created     time.Time
	Modified    time.Time
}

func NewLink(title string, url string, desc string, tags []string, when time.Time) *Link {
	e := new(Link)

	// User supplied content
	e.Title = title
	e.Url = url
	e.Description = desc
	e.Tags = tags
	e.Posted = when

	// Computer generated content
	e.Created = time.Now()
	e.Modified = time.Now()

	return e
}

func (l *Link) TagString() string {
	tags := []string{}
	for _, t := range l.Tags {
		tags = append(tags, fmt.Sprintf("#%s", t))
	}

	return strings.Join(tags, " ")
}

func (e *Link) Save(c context.Context) error {
	k := datastore.NewKey(c, "Link", e.Url, 0, nil)
	k2, err := datastore.Put(c, k, e)
	if err == nil {
		log.Infof(c, "Wrote %+v", e)
		log.Infof(c, "Old key: %+v; New Key: %+v", k, k2)
	} else {
		log.Warningf(c, "Error writing link: %v", e)
	}
	return err
}

func AllLinks(c context.Context) (*[]Link, error) {
	return Links(c, -1, true)
}

func Links(c context.Context, limit int, recentFirst bool) (*[]Link, error) {
	q := datastore.NewQuery("Link")

	if recentFirst {
		q = q.Order("-Posted")
	} else {
		q = q.Order("Posted")
	}

	if limit > 0 {
		q = q.Limit(limit)
	}

	links := new([]Link)
	_, err := q.GetAll(c, links)
	return links, err
}
