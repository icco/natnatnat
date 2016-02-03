package main

import (
	"net/http"
	"strconv"

	"github.com/pilu/traffic"
	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
	"google.golang.org/appengine/search"
	"google.golang.org/appengine/user"
)

// https://cloud.google.com/appengine/docs/go/search/query_strings
// https://godoc.org/google.golang.org/appengine/search
// https://cloud.google.com/appengine/docs/go/search

type SearchData struct {
	Count   int
	IsAdmin bool
}

func SearchHandler(w traffic.ResponseWriter, r *traffic.Request) {
	c := appengine.NewContext(r.Request)
	s_val := r.Request.FormValue("s")

	count := 0
	if s_val != "" {
		index, err := search.Open("entries")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		iter := index.Search(c, s_val, nil)
		count = iter.Count()
		log.Infof(c, "Search Results for %+v: %+v", s_val, iter)
	}

	w.Render("search", &SearchData{
		Count:   count,
		IsAdmin: user.IsAdmin(c),
	})
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
		log.Errorf(c, err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	for _, v := range *entries {
		_, err = index.Put(c, strconv.FormatInt(v.Id, 10), v.SearchDoc())
		log.Infof(c, "Trying to put %+v: %+v", v.Id, strconv.FormatInt(v.Id, 10))
		if err != nil {
			log.Errorf(c, err.Error())
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}
