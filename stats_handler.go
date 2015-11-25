package main

import (
	"net/http"
	"strings"
	"time"

	"google.golang.org/appengine"
	"google.golang.org/appengine/user"

	"github.com/pilu/traffic"
)

type StatsData struct {
	Posts        int
	WordsPerPost float64
	PostsPerDay  float64
	WordsPerDay  float64
	DaysSince    float64
	IsAdmin      bool
	LinksPerPost float64
	LinksPerDay  float64
}

func StatsHandler(w traffic.ResponseWriter, r *traffic.Request) {
	c := appengine.NewContext(r.Request)
	entries, err := AllPosts(c)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	postCount := len(*entries)
	oldestPost := (*entries)[postCount-1]
	dayCount := time.Since(oldestPost.Datetime).Hours() / 24.0

	words := 0
	for _, p := range *entries {
		words += len(strings.Fields(p.Content))
		words += len(strings.Fields(p.Title))
	}

	links, err := AllLinks(c)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	readLinks := len(*links)

	data := &StatsData{
		Posts:        postCount,
		PostsPerDay:  float64(postCount) / dayCount,
		WordsPerPost: float64(words) / float64(postCount),
		WordsPerDay:  float64(words) / dayCount,
		DaysSince:    dayCount,
		IsAdmin:      user.IsAdmin(c),
		LinksPerDay:  float64(readLinks) / dayCount,
	}
	w.Render("stats", data)
}

func StatsHistoryJsonHandler(w traffic.ResponseWriter, r *traffic.Request) {
	c := appengine.NewContext(r.Request)
	entries, err := AllPosts(c)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteJSON(entries)
}
