package main

import (
	"net/http"

	"github.com/pilu/traffic"
	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
	"google.golang.org/appengine/search"
)

// https://cloud.google.com/appengine/docs/go/search/query_strings
// https://godoc.org/google.golang.org/appengine/search
// https://cloud.google.com/appengine/docs/go/search

func SearchHandler(w traffic.ResponseWriter, r *traffic.Request) {
	c := appengine.NewContext(r.Request)
	s_val := r.Request.FormValue("s")

	if s_val != "" {
		index, err := search.Open("entries")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		iter := index.Search(c, s_val, nil)
		log.Infof(c, "Search Results: %+v", iter)
	}

	w.Render("search", nil)
}

func SearchWorkHandler(w traffic.ResponseWriter, r *traffic.Request) {
	c := appengine.NewContext(r.Request)
	entries, err := AllPosts(c)
	if err != nil {
		log.Errorf(c, err.Error())
		http.Error(w, err.Error(), 500)
		return
	}

	index, err := search.Open("entries")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	for _, v := range *entries {
		_, err = index.Put(c, string(v.Id), v)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}
