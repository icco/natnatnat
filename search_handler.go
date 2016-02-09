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
	Results *[]Entry
	IsAdmin bool
}

func SearchHandler(w traffic.ResponseWriter, r *traffic.Request) {
	c := appengine.NewContext(r.Request)
	s_val := r.Request.FormValue("s")

	results := []Entry{}

	if s_val != "" {
		index, err := search.Open("entries")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		iter := index.Search(c, s_val, nil)
		log.Infof(c, "Search Results for %+v: %+v", s_val, iter)

		for t := iter; ; {
			var doc EntrySearch
			id, err := t.Next(&doc)
			if err == search.Done {
				break
			}
			if err != nil {
				log.Errorf(c, "Error iterating: %+v", err)
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			log.Debugf(c, "%s -> %#v\n", id, doc)

			entry, err := GetEntry(c, int64(doc.Id))
			if err != nil {
				log.Errorf(c, "Error getting entry: %+v", err)
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			results = append(results, *entry)
		}
	}

	w.Render("search", &SearchData{
		Count:   len(results),
		IsAdmin: user.IsAdmin(c),
		Results: &results,
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
