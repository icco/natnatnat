package main

import (
	"fmt"
	"html/template"
	"net/http"
	"time"

	"github.com/gorilla/feeds"
	"github.com/pilu/traffic"
	"golang.org/x/net/context"
	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
)

var baseUrl = "https://writing.natwelch.com"

func buildFeed(c context.Context, entries *[]Entry) *feeds.Feed {
	now := time.Now()
	me := &feeds.Author{"Nat Welch", "nat@natwelch.com"}
	feed := &feeds.Feed{
		Title:       "Nat? Nat. Nat!",
		Link:        &feeds.Link{Href: baseUrl},
		Description: "Thoughts from Nat about stuff",
		Author:      me,
		Created:     now,
	}

	feed.Items = []*feeds.Item{}
	for _, v := range *entries {
		if !v.Draft {
			title := fmt.Sprintf("Nat? Nat. Nat! #%d", v.Id)
			if v.Title != "" {
				title = v.Title
			}

			feed.Items = append(feed.Items, &feeds.Item{
				Title:       title,
				Link:        &feeds.Link{Href: baseUrl + v.Url()},
				Description: string(v.Html()),
				Author:      me,
				Created:     v.Datetime,
				Updated:     v.Modified,
			})
		}
	}

	return feed
}

func FeedAtomHandler(w traffic.ResponseWriter, r *traffic.Request) {
	c := appengine.NewContext(r.Request)
	entries, err := RecentPosts(c)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	feed := buildFeed(c, entries)
	atom, err := feed.ToAtom()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.Header().Set("Content-Type", "application/atom+xml; charset=utf-8")
	w.WriteText("%s", atom)
	return
}

func FeedRssHandler(w traffic.ResponseWriter, r *traffic.Request) {
	c := appengine.NewContext(r.Request)
	entries, err := RecentPosts(c)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	feed := buildFeed(c, entries)
	rss, err := feed.ToRss()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.Header().Set("Content-Type", "application/rss+xml; charset=utf-8")
	w.WriteText("%s", rss)
	return
}

func SummaryAtomHandler(w traffic.ResponseWriter, r *traffic.Request) {
	http.Redirect(w, r.Request, "/feed.atom", 301)
}

func SummaryRssHandler(w traffic.ResponseWriter, r *traffic.Request) {
	http.Redirect(w, r.Request, "/feed.rss", 301)
}

type SummaryJson struct {
	Id       int64         `json:"id"`
	Title    string        `json:"title"`
	Html     template.HTML `json:"html"`
	Datetime time.Time     `json:"date"`
	Tags     []string      `json:"tags"`
}

func SummaryJsonHandler(w traffic.ResponseWriter, r *traffic.Request) {
	c := appengine.NewContext(r.Request)
	entries, err := AllPosts(c)
	if err != nil {
		log.Errorf(c, "Error loading posts: %+v", err)
		http.Error(w, err.Error(), 500)
		return
	}
	posts := []SummaryJson{}
	for _, entry := range *entries {
		post := &SummaryJson{}
		post.Id = entry.Id
		post.Title = entry.Title
		post.Html = Markdown(SummarizeText(entry.Content))
		post.Datetime = entry.Datetime
		post.Tags = entry.Tags
		posts = append(posts, *post)
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteJSON(posts)
}
