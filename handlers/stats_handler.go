package handlers

import (
	"net/http"
	"strings"
	"time"

	"appengine"
	"appengine/user"

	"github.com/icco/natnatnat/models"
	"github.com/pilu/traffic"
)

type StatsData struct {
	Posts        int
	WordsPerPost float64
	PostsPerDay  float64
	IsAdmin      bool
}

func StatsHandler(w traffic.ResponseWriter, r *traffic.Request) {
	c := appengine.NewContext(r.Request)
	entries, err := models.AllPosts(c)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	postCount := len(*entries)
	oldestPost := (*entries)[postCount-1]
	dayCount := time.Since(oldestPost.Datetime).Hours() / 24

	words := 0
	for _, p := range *entries {
		words += len(strings.Fields(p.Content))
		words += len(strings.Fields(p.Title))
	}

	data := &StatsData{
		Posts:        postCount,
		PostsPerDay:  float64(postCount) / dayCount,
		WordsPerPost: float64(words) / float64(postCount),
		WordsPerDay:  float64(words) / dayCount,
		IsAdmin:      user.IsAdmin(c),
	}
	w.Render("stats", data)
}

func StatsHistoryJsonHandler(w traffic.ResponseWriter, r *traffic.Request) {
	c := appengine.NewContext(r.Request)
	entries, err := models.AllPosts(c)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.WriteJSON(entries)
}
