package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/pilu/traffic"
	"github.com/spf13/hugo/parser"
	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
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
	c := appengine.NewContext(r.Request)

	// Read drafts from disk
	dir := "./longform/drafts/"
	drafts, err := ioutil.ReadDir(dir)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	log.Warningf(c, "Drafts: %+v", drafts)

	// Iterate through
	for _, file := range drafts {
		if strings.HasPrefix(file.Name(), "20") && file.Mode().IsRegular() {
			log.Infof(c, "Draft: %+v", file.Name())
			f, err := os.Open(dir + file.Name())
			defer f.Close()
			if err != nil {
				log.Errorf(c, "Error reading file: %v", err.Error())
				http.Error(w, err.Error(), 500)
				return
			}
			// get the page from file
			p, err := parser.ReadFrom(f)
			if err != nil {
				log.Errorf(c, "Error reading file: %v", err.Error())
				http.Error(w, err.Error(), 500)
				return
			}
			meta, err := p.Metadata()
			if err != nil {
				log.Errorf(c, "Error reading file: %v", err.Error())
				http.Error(w, err.Error(), 500)
				return
			}

			// create entries for those that don't exist, update those that do
			log.Debugf(c, "Opened File: %+v", meta)
		}
	}

	// Read posts from disk,
	dir = "./longform/posts/"
	posts, err := ioutil.ReadDir(dir)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// create entries for those that don't exist, update those that do
	for _, file := range posts {
		if strings.HasPrefix(file.Name(), "20") && file.Mode().IsRegular() {
			log.Infof(c, "Post: %+v", file.Name())
		}
	}
}
