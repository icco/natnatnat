package handlers

import (
	"appengine"
	"appengine/user"
	"github.com/icco/natnatnat/models"
	"github.com/pilu/traffic"
	"net/http"
)

type RootData struct {
	Posts   interface{}
	IsAdmin bool
}

func RootHandler(w traffic.ResponseWriter, r *traffic.Request) {
	c := appengine.NewContext(r.Request)
	entries, err := models.AllPosts(c)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	data := &RootData{Posts: entries, IsAdmin: user.IsAdmin(c)}
	w.Render("index", data)
}
