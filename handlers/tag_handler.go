package handlers

import (
	"net/http"

	"appengine"

	"github.com/icco/natnatnat/models"
	"github.com/pilu/traffic"
)

type TagData struct {
	Posts interface{}
	Tag   string
}

func TagHandler(w traffic.ResponseWriter, r *traffic.Request) {
	c := appengine.NewContext(r.Request)
	tag := r.Param("id")

	if tag == "" {
		http.Redirect(w, r.Request, "/tags", 301)
	}

	entries, err := models.PostsWithTag(c, tag)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	data := &TagData{Posts: entries, Tag: tag}
	w.Render("tag", data)
}

func TagsHandler(w traffic.ResponseWriter, r *traffic.Request) {
	http.Error(w, "Sorry, I haven't implemented this yet", 500)
}
