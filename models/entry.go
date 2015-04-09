package models

import (
	"fmt"
	"html/template"
	"regexp"
	"strings"
	"time"

	"appengine"
	"appengine/datastore"

	"github.com/kennygrant/sanitize"
	"github.com/russross/blackfriday"
)

type Entry struct {
	Id       int64     `json:"id"`
	Title    string    `json:"title"`         // optional
	Content  string    `datastore:",noindex"` // Markdown
	Datetime time.Time `json:"date"`
	Created  time.Time
	Modified time.Time
	Tags     []string `json:"tags"`
	Public   bool     `json:"page"`
}

var HashtagRegex *regexp.Regexp = regexp.MustCompile(`(\s)#(\w+)`)
var TwitterHandleRegex *regexp.Regexp = regexp.MustCompile(`(\s)@([_A-Za-z0-9]+)`)

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

func GetEntry(c appengine.Context, id int64) (*Entry, error) {
	var entry Entry
	q := datastore.NewQuery("Entry").Filter("Id =", id).Limit(1)
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
	return Posts(c, -1, true)
}

func Posts(c appengine.Context, limit int, recentFirst bool) (*[]Entry, error) {
	q := datastore.NewQuery("Entry").Filter("Public =", true)

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

func RecentPosts(c appengine.Context) (*[]Entry, error) {
	return Posts(c, 20, true)
}

func (e *Entry) HasId() bool {
	return (e.Id > 0)
}

func (e *Entry) Save(c appengine.Context) error {
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
		c.Infof("Wrote %+v", e)
		c.Infof("Old key: %+v; New Key: %+v", k, k2)
	} else {
		c.Warningf("Error writing entry: %v", e)
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

func (e *Entry) PrevPost(c appengine.Context) string {
	var entry Entry
	q := datastore.NewQuery("Entry").Order("-Datetime").Filter("Datetime <", e.Datetime).Limit(1)
	_, err := q.Run(c).Next(&entry)
	if err != nil {
		c.Infof("Error getting previous post for %d.", e.Id)
		return ""
	}

	return entry.Url()
}

func (e *Entry) NextPost(c appengine.Context) string {
	var entry Entry
	q := datastore.NewQuery("Entry").Order("Datetime").Filter("Datetime >", e.Datetime).Limit(1)
	_, err := q.Run(c).Next(&entry)
	if err != nil {
		c.Infof("Error getting next post for %d.", e.Id)
		return ""
	}

	return entry.Url()
}

// TODO(icco): Actually finish this.
func GetLinksFromContent(c appengine.Context, content string) ([]string, error) {
	httpRegex := regexp.MustCompile(`http:\/\/((\w|\.)+)`)
	matches := httpRegex.FindAllString(content, -1)
	if matches == nil {
		return []string{}, nil
	}

	for _, match := range matches {
		c.Infof("%+v", match)
	}

	return []string{}, nil
}

func PostsWithTag(c appengine.Context, tag string) (*[]Entry, error) {
	q := datastore.NewQuery("Entry").Order("-Datetime").Filter("Tags =", tag)
	entries := new([]Entry)
	_, err := q.GetAll(c, entries)
	return entries, err
}

// Markdown.
func Markdown(args ...interface{}) template.HTML {
	inc := []byte(fmt.Sprintf("%s", args...))
	inc = twitterHandleToMarkdown(inc)
	inc = hashTagsToMarkdown(inc)
	s := blackfriday.MarkdownCommon(inc)
	return template.HTML(s)
}

func twitterHandleToMarkdown(in []byte) []byte {
	return TwitterHandleRegex.ReplaceAll(in, []byte("$1[@$2](http://twitter.com/$2)"))
}

func hashTagsToMarkdown(in []byte) []byte {
	return HashtagRegex.ReplaceAll(in, []byte("$1[#$2](/tags/$2)"))
}
