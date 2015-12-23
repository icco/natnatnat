package main

import (
	"fmt"
	"sort"
	"strings"
	"time"

	"golang.org/x/net/context"

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

// Data for building a list of links by day.
type LinkDay struct {
	Links []Link
	Day   time.Time
}

type linkDays []*LinkDay

// These three functions are needed for Sort.
func (p linkDays) Len() int           { return len(p) }
func (p linkDays) Less(i, j int) bool { return p[i].Day.Before(p[j].Day) }
func (p linkDays) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p linkDays) HasDate(day time.Time) int {
	for i, d := range p {
		if d.Day == day {
			return i
		}
	}
	return -1
}

func LinksByDay(c context.Context, days int) (*linkDays, error) {
	links, err := AllLinks(c)
	if err != nil {
		return nil, err
	}

	ds := make([]*LinkDay, 0)

	for _, l := range *links {
		date := l.Posted.Round(time.Hour * 24)
		if linkDays(ds).HasDate(date) < 0 {
			ds = append(ds, &LinkDay{
				Day:   date,
				Links: []Link{l},
			})
		} else {
			i := linkDays(ds).HasDate(date)
			ds[i].Links = append(ds[i].Links, l)
		}
	}

	lds := linkDays(ds)
	sort.Reverse(lds)

	subset := lds[0:days]

	return &subset, nil
}
