package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/pilu/traffic"
	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
	"google.golang.org/appengine/taskqueue"
	//"gopkg.in/yaml.v2"
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
	c := appengine.NewContext(r.Request)

	// Read drafts from disk
	drafts, err := ioutil.ReadDir("./longform/drafts")
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	log.Warningf(c, "Drafts: %+v", drafts)

	// Iterate through, create entries for those that don't exist, update those that do
	for _, file := range drafts {
		if strings.HasPrefix(file.Name(), "20") {
			log.Infof(c, "Draft: %+v", file.Name())
		}
	}

	// Read posts from disk,
	posts, err := ioutil.ReadDir("./longform/posts")
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// create entries for those that don't exist, update those that do
	for _, file := range posts {
		if strings.HasPrefix(file.Name(), "20") {
			log.Infof(c, "Post: %+v", file.Name())
		}
	}
}
