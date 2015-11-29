package main

import (
	"fmt"
	"html/template"
	"regexp"
	"strings"
	"time"

	"google.golang.org/appengine/datastore"
	"google.golang.org/appengine/log"

	"golang.org/x/net/context"

	"github.com/kennygrant/sanitize"
)

type Entry struct {
	Id       int64     `json:"id"`
	Title    string    `json:"title"`                     // optional
	Content  string    `datastore:",noindex" json:"text"` // Markdown
	Datetime time.Time `json:"date"`
	Created  time.Time `json:"created"`
	Modified time.Time `json:"modified"`
	Tags     []string  `json:"tags"`
	Public   bool      `json:"-"`
	Longform string    `json:"-"`
	Draft    bool      `json:"-"`
}

func NewEntry(title string, content string, datetime time.Time, public bool, tags []string) *Entry {
	e := new(Entry)

	// User supplied content
	e.Title = title
	e.Content = content
	e.Datetime = datetime
	e.Tags = tags
	e.Public = public

	// Computer generated content
	e.Created = time.Now()
	e.Modified = time.Now()
	e.Draft = false

	return e
}

func ParseTags(text string) ([]string, error) {
	// http://golang.org/pkg/regexp/#Regexp.FindAllStringSubmatch
	finds := HashtagRegex.FindAllStringSubmatch(text, -1)
	ret := make([]string, 0)
	for _, v := range finds {
		if len(v) > 2 {
			ret = append(ret, strings.ToLower(v[2]))
		}
	}

	return ret, nil
}

func GetEntry(c context.Context, id int64) (*Entry, error) {
	var entry Entry
	q := datastore.NewQuery("Entry").Filter("Id =", id).Limit(1)
	_, err := q.Run(c).Next(&entry)
	if err != nil {
		return nil, err
	}

	return &entry, nil
}

func GetLongform(c context.Context, longform string) (*Entry, error) {
	var entry Entry
	q := datastore.NewQuery("Entry").Filter("Longform =", longform).Limit(1)
	_, err := q.Run(c).Next(&entry)
	if err != nil {
		return nil, err
	}

	return &entry, nil
}

func MaxId(c context.Context) (int64, error) {
	var entry Entry
	q := datastore.NewQuery("Entry").Order("-Id").Limit(1)
	_, err := q.Run(c).Next(&entry)
	if err != nil {
		return 0, err
	}
	return entry.Id, nil
}

func AllPosts(c context.Context) (*[]Entry, error) {
	return Posts(c, -1, true)
}

func Pagination(c context.Context, posts, offset int) (*[]Entry, error) {
	q := datastore.NewQuery("Entry").
		Filter("Public =", true).
		Order("-Datetime").
		Limit(posts).
		Offset(offset)

	entries := new([]Entry)
	_, err := q.GetAll(c, entries)
	return entries, err
}

func ArchivePageQuery() *datastore.Query {
	return datastore.NewQuery("Entry").
		Filter("Public =", true).
		Project("Id", "Datetime").
		Order("-Datetime").
		Limit(50)
}

func Posts(c context.Context, limit int, recentFirst bool) (*[]Entry, error) {
	q := datastore.NewQuery("Entry").Filter("Public =", true).Filter("Draft ==", false)

	if recentFirst {
		q = q.Order("-Datetime")
	} else {
		q = q.Order("Datetime")
	}

	if limit > 0 {
		q = q.Limit(limit)
	}

	entries := new([]Entry)
	_, err := q.GetAll(c, entries)
	return entries, err
}

func Drafts(c context.Context) (*[]Entry, error) {
	q := datastore.NewQuery("Entry").Filter("Draft =", true)
	q = q.Order("-Datetime")
	entries := new([]Entry)
	_, err := q.GetAll(c, entries)
	return entries, err
}

func RecentPosts(c context.Context) (*[]Entry, error) {
	return Posts(c, 20, true)
}

func (e *Entry) HasId() bool {
	return (e.Id > 0)
}

func (e *Entry) Save(c context.Context) error {
	var k *datastore.Key
	if !e.HasId() {
		id, _ := MaxId(c)
		e.Id = id + 1
		k = datastore.NewIncompleteKey(c, "Entry", nil)
	} else {
		// Find the key
		var err error
		q := datastore.NewQuery("Entry").Filter("Id =", e.Id).Limit(1).KeysOnly()
		k, err = q.Run(c).Next(nil)
		if err != nil {
			return err
		}
	}

	// Pull out links
	GetLinksFromContent(c, e.Content)

	// Figure out Tags
	tags, err := ParseTags(e.Content)
	if err != nil {
		return err
	}
	e.Tags = tags

	k2, err := datastore.Put(c, k, e)
	if err == nil {
		log.Infof(c, "Wrote %+v", e)
		log.Infof(c, "Old key: %+v; New Key: %+v", k, k2)
	} else {
		log.Warningf(c, "Error writing entry: %v", e)
	}
	return err
}

func (e *Entry) Url() string {
	return fmt.Sprintf("/post/%d", e.Id)
}

func (e *Entry) EditUrl() string {
	return fmt.Sprintf("/edit/%d", e.Id)
}

func (e *Entry) Html() template.HTML {
	return Markdown(e.Content)
}

func (e *Entry) Summary() string {
	// truncate(strip_tags(m(p.text)), :length => 100).strip
	stripped := sanitize.HTML(string(e.Html()))
	if len(stripped) > 100 {
		return fmt.Sprintf("%s...", stripped[:100])
	} else {
		return stripped
	}
}

func (e *Entry) PrevPost(c context.Context) string {
	var entry Entry
	q := datastore.NewQuery("Entry").Order("-Datetime").Filter("Datetime <", e.Datetime).Limit(1)
	_, err := q.Run(c).Next(&entry)
	if err != nil {
		log.Infof(c, "Error getting previous post for %d.", e.Id)
		return ""
	}

	return entry.Url()
}

func (e *Entry) NextPost(c context.Context) string {
	var entry Entry
	q := datastore.NewQuery("Entry").Order("Datetime").Filter("Datetime >", e.Datetime).Limit(1)
	_, err := q.Run(c).Next(&entry)
	if err != nil {
		log.Infof(c, "Error getting next post for %d.", e.Id)
		return ""
	}

	return entry.Url()
}

// TODO(icco): Actually finish this.
func GetLinksFromContent(c context.Context, content string) ([]string, error) {
	httpRegex := regexp.MustCompile(`http:\/\/((\w|\.)+)`)
	matches := httpRegex.FindAllString(content, -1)
	if matches == nil {
		return []string{}, nil
	}

	for _, match := range matches {
		log.Infof(c, "%+v", match)
	}

	return []string{}, nil
}

func PostsForDay(c context.Context, year, month, day int64) (*[]Entry, error) {
	entries := new([]Entry)
	start := time.Date(int(year), time.Month(month), int(day), 0, 0, 0, 0, time.UTC)
	end := start.AddDate(0, 0, 1)
	q := datastore.NewQuery("Entry").Order("-Datetime").Filter("Datetime >=", start).Filter("Datetime <", end).Filter("Public =", true)
	_, err := q.GetAll(c, entries)
	return entries, err
}

func PostsWithTag(c context.Context, tag string) (*map[int64]Entry, error) {
	entries := make(map[int64]Entry, 0)
	aliases := GetTagAliases(c, tag)
	aliasesAndTag := append(*aliases, tag)

	for _, v := range aliasesAndTag {
		more_entries := new([]Entry)
		q := datastore.NewQuery("Entry").Order("-Datetime").Filter("Tags =", v)
		_, err := q.GetAll(c, more_entries)
		if err != nil {
			return &entries, err
		}
		for _, e := range *more_entries {
			entries[e.Id] = e
		}
	}

	return &entries, nil
}
