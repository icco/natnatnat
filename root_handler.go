package main

import (
	"appengine"
	"github.com/pilu/traffic"
	"net/http"
)

type RootData struct {
	Posts interface{}
}

func RootHandler(w traffic.ResponseWriter, r *traffic.Request) {
	c := appengine.NewContext(r.Request)
	entries, err := AllPosts(c)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	data := &RootData{entries}
	w.Render("index", data)
}
