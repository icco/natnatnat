package main

import (
	"fmt"
	"html/template"
	"math"
	"regexp"
	"strings"
	"time"

	"github.com/kennygrant/sanitize"
	"golang.org/x/net/context"
	"google.golang.org/appengine/datastore"
	"google.golang.org/appengine/log"
	"google.golang.org/appengine/search"
)

type Entry struct {
	Id       int64     `json:"id"`
	Title    string    `json:"title"`                     // optional
	Content  string    `datastore:",noindex" json:"text"` // Markdown
	Datetime time.Time `json:"date"`
	Created  time.Time `json:"created"`
	Modified time.Time `json:"modified"`
	Tags     []string  `json:"tags"`
	Longform string    `json:"-"`
	Public   bool      `json:"-"` // Deprecated
	Draft    bool      `json:"-"`
}

type EntryJson struct {
	Entry
	Html     template.HTML `json:"html"`
	ReadTime int           `json:"readtime"`
}

type EntrySearch struct {
	Id       float64
	Title    string
	Content  search.HTML
	Datetime time.Time
	Tags     string
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

	log.Infof(c, "Attempted to find %v. %+v %+v", longform, entry, err)

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

func AllPostsJson(c context.Context) (*[]EntryJson, error) {
	posts := []EntryJson{}
	entries, err := AllPosts(c)
	if err != nil {
		return nil, err
	}
	for _, entry := range *entries {
		post := &EntryJson{
			Html:     entry.Html(),
			ReadTime: entry.ReadTime(),
		}
		post.Id = entry.Id
		post.Title = entry.Title
		post.Content = entry.Content
		post.Datetime = entry.Datetime
		post.Created = entry.Created
		post.Modified = entry.Modified
		post.Tags = entry.Tags
		posts = append(posts, *post)
	}
	return &posts, nil
}

func Pagination(c context.Context, posts, offset int) (*[]Entry, error) {
	q := datastore.NewQuery("Entry").
		Filter("Draft =", false).
		Order("-Datetime").
		Limit(posts).
		Offset(offset)

	entries := new([]Entry)
	_, err := q.GetAll(c, entries)
	return entries, err
}

func ArchivePageQuery() *datastore.Query {
	return datastore.NewQuery("Entry").
		Filter("Draft =", false).
		Project("Id", "Datetime").
		Order("-Datetime").
		Limit(50)
}

func Posts(c context.Context, limit int, recentFirst bool) (*[]Entry, error) {
	q := datastore.NewQuery("Entry").Filter("Draft =", false)

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
	q := datastore.NewQuery("Entry").Filter("Draft =", true).Order("-Datetime")
	entries := new([]Entry)
	_, err := q.GetAll(c, entries)
	return entries, err
}

func LongformPosts(c context.Context) (*[]Entry, error) {
	q := datastore.NewQuery("Entry").Filter("Longform >", "").Order("-Longform")
	entries := new([]Entry)
	_, err := q.GetAll(c, entries)
	return entries, err
}

func RecentPosts(c context.Context) (*[]Entry, error) {
	return Posts(c, 20, true)
}

func (e *Entry) SearchDoc() *EntrySearch {
	return &EntrySearch{
		Id:       float64(e.Id),
		Title:    e.Title,
		Content:  search.HTML(e.Html()),
		Datetime: e.Datetime,
		Tags:     strings.Join(e.Tags, ","),
	}
}

func (e *Entry) HasId() bool {
	return (e.Id > 0)
}

func (e *Entry) Delete(c context.Context) error {
	// Find the key
	var err error
	q := datastore.NewQuery("Entry").Filter("Id =", e.Id).Limit(1).KeysOnly()
	k, err := q.Run(c).Next(nil)
	if err != nil {
		return err
	}

	return datastore.Delete(c, k)
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

	cnt, err := datastore.NewQuery("Entry").Filter("Id =", e.Id).Count(c)
	if err != nil {
		return err
	}
	log.Infof(c, "ID: %v: %v", e.Id, cnt)
	if cnt >= 2 {
		id, _ := MaxId(c)
		e.Id = id + 1
	}

	// Pull out links
	// TODO: Do something with the output
	GetLinksFromContent(c, e.Content)

	// Figure out Tags
	tags, err := ParseTags(e.Content)
	if err != nil {
		return err
	}
	e.Tags = tags

	k2, err := datastore.Put(c, k, e)
	if err == nil {
		// log.Infof(c, "Wrote %+v", e)
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

func (e *Entry) ReadTime() int {
	ReadingSpeed := 265.0
	words := len(strings.Split(e.Content, " "))
	seconds := int(math.Ceil(float64(words) / ReadingSpeed * 60.0))

	return seconds
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
	q := datastore.NewQuery("Entry").Order("-Datetime").Filter("Datetime <", e.Datetime).Filter("Draft =", false).Limit(1)
	_, err := q.Run(c).Next(&entry)
	if err != nil {
		log.Infof(c, "Error getting previous post for %d.", e.Id)
		return ""
	}

	return entry.Url()
}

func (e *Entry) NextPost(c context.Context) string {
	var entry Entry
	q := datastore.NewQuery("Entry").Order("Datetime").Filter("Datetime >", e.Datetime).Filter("Draft =", false).Limit(1)
	_, err := q.Run(c).Next(&entry)
	if err != nil {
		log.Infof(c, "Error getting next post for %d.", e.Id)
		return ""
	}

	return entry.Url()
}

func GetLinksFromContent(c context.Context, content string) ([]string, error) {
	httpRegex := regexp.MustCompile(`https?:\/\/((\w|\.)+)`)
	matches := httpRegex.FindAllString(content, -1)
	if matches == nil {
		return []string{}, nil
	}

	log.Infof(c, "URLs Found: %+v", matches)

	return matches, nil
}

func PostsForDay(c context.Context, year, month, day int64) (*[]Entry, error) {
	entries := new([]Entry)
	start := time.Date(int(year), time.Month(month), int(day), 0, 0, 0, 0, time.UTC)
	end := start.AddDate(0, 0, 1)
	q := datastore.NewQuery("Entry").Order("-Datetime").Filter("Datetime >=", start).Filter("Datetime <", end).Filter("Draft =", false)
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
