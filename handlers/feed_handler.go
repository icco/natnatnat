package handlers

import (
	"fmt"
	"net/http"
	"time"

	"appengine"

	"github.com/gorilla/feeds"
	"github.com/icco/natnatnat/models"
	"github.com/pilu/traffic"
)

func buildFeed(c appengine.Context, entries *[]models.Entry) *feeds.Feed {
	now := time.Now()
	me := &feeds.Author{"Nat Welch", "nat@natwelch.com"}
	feed := &feeds.Feed{
		Title:       "Nat? Nat. Nat!",
		Link:        &feeds.Link{Href: "https://writing.natwelch.com"},
		Description: "Thoughts from Nat about stuff",
		Author:      me,
		Created:     now,
	}

	feed.Items = []*feeds.Item{}
	for _, v := range *entries {
		if v.Public {
			feed.Items = append(feed.Items, &feeds.Item{
				Title:       v.Title,
				Link:        &feeds.Link{Href: v.Url()},
				Description: string(v.Html()),
				Author:      me,
				Created:     v.Datetime,
				Updated:     v.Modified,
			})
		}
	}

	return feed
}

func buildSummary(c appengine.Context, entries *[]models.Entry) *feeds.Feed {
	now := time.Now()
	me := &feeds.Author{"Nat Welch", "nat@natwelch.com"}
	feed := &feeds.Feed{
		Title:       "Nat? Nat. Nat! - Summaries",
		Link:        &feeds.Link{Href: "https://writing.natwelch.com"},
		Description: "Thoughts from Nat about stuff",
		Author:      me,
		Created:     now,
	}

	feed.Items = []*feeds.Item{}
	for _, v := range *entries {
		if v.Public {
			feed.Items = append(feed.Items, &feeds.Item{
				Title:       fmt.Sprintf("Nat? Nat. Nat! #%d", v.Id),
				Link:        &feeds.Link{Href: v.Url()},
				Description: "At some point, I'll put something here.",
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
	entries, err := models.AllPosts(c)
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
	entries, err := models.AllPosts(c)
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
	c := appengine.NewContext(r.Request)
	entries, err := models.AllPosts(c)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	feed := buildSummary(c, entries)
	atom, err := feed.ToAtom()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.Header().Set("Content-Type", "application/atom+xml; charset=utf-8")
	w.WriteText("%s", atom)
	return
}

func SummaryRssHandler(w traffic.ResponseWriter, r *traffic.Request) {
	c := appengine.NewContext(r.Request)
	entries, err := models.AllPosts(c)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	feed := buildSummary(c, entries)
	rss, err := feed.ToRss()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.Header().Set("Content-Type", "application/rss+xml; charset=utf-8")
	w.WriteText("%s", rss)
	return
}
