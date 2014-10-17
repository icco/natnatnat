package handlers

import (
	"appengine"
	"fmt"
	"github.com/gorilla/feeds"
	"github.com/icco/natnatnat/models"
	"github.com/pilu/traffic"
	"time"
)

func buildFeed(c appengine.Context, entries *[]models.Entry) feeds.Feed {
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
	for _, v := range entries {
		feed.Items = append(feed.Items, &feeds.Item{
			Title:       v.Title,
			Link:        &feeds.Link{Href: v.Url()},
			Description: v.Html(),
			Author:      me,
			Created:     v.Datetime,
			Updated:     v.Updated,
		})
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
	atom, err := feed.ToAtom()
}

func FeedRssHandler(w traffic.ResponseWriter, r *traffic.Request) {
	rss, err := feed.ToRss()
}
