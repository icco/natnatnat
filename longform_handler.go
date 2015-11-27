package main

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/pilu/traffic"

	"google.golang.org/appengine"
	"google.golang.org/appengine/taskqueue"
)

func LongformQueueHandler(w traffic.ResponseWriter, r *traffic.Request) {
	c := appengine.NewContext(r.Request)
	t := taskqueue.NewPOSTTask("/longform/work", url.Values{})
	_, err := taskqueue.Add(c, t, "")

	if err != nil {
		http.Error(w, err.Error(), 500)
	} else {
		fmt.Fprint(w, "success.\n")
	}
}

func LongformWorkHandler(w traffic.ResponseWriter, r *traffic.Request) {
	// c := appengine.NewContext(r.Request)
	// Stub:
	// Read drafts from disk
	// Iterate through, create entries for those that don't exist, update those that do
	// Read posts from disk, create entries for those that don't exist, update those that do
}
