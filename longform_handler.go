package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/pilu/traffic"
	"github.com/spf13/cast"
	"github.com/spf13/hugo/parser"
	"golang.org/x/net/context"
	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
	"google.golang.org/appengine/log"
)

// For old urls still hosted in github pages.
func PseudowebHandler(w traffic.ResponseWriter, r *traffic.Request) {
	http.Redirect(w, r.Request, fmt.Sprintf("http://pseudoweb.net%s", r.Request.URL.Path), 301)
}

func LongformJsonHandler(w traffic.ResponseWriter, r *traffic.Request) {
	c := appengine.NewContext(r.Request)
	entries, err := LongformPosts(c)
	if err != nil {
		log.Errorf(c, "Error getting posts: %v", err.Error())
		http.Error(w, err.Error(), 500)
		return
	}

	data := map[string]int64{}
	for _, e := range *entries {
		data[e.Longform] = e.Id
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteJSON(data)
}

func LongformWorkHandler(w traffic.ResponseWriter, r *traffic.Request) {
	c := appengine.NewContext(r.Request)

	// Read drafts from disk
	dir := "./longform/drafts/"
	drafts, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Errorf(c, "Error opening directory: %v", err.Error())
		http.Error(w, err.Error(), 500)
		return
	}

	// Iterate through
	for _, file := range drafts {
		if strings.HasPrefix(file.Name(), "20") && file.Mode().IsRegular() {
			draft := true
			err := createPostFromLongformFile(c, dir, file, draft)
			if err != nil {
				http.Error(w, err.Error(), 500)
				return
			}
		}
	}

	// Read posts from disk,
	dir = "./longform/posts/"
	posts, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Errorf(c, "Error opening directory: %v", err.Error())
		http.Error(w, err.Error(), 500)
		return
	}

	// create entries for those that don't exist, update those that do
	for _, file := range posts {
		if strings.HasPrefix(file.Name(), "20") && file.Mode().IsRegular() {
			draft := false
			err := createPostFromLongformFile(c, dir, file, draft)
			if err != nil {
				http.Error(w, err.Error(), 500)
				return
			}
		}
	}

	all_posts, err := LongformPosts(c)
	for _, e := range *all_posts {
		dir := "./longform/posts/"
		if e.Draft {
			dir = "./longform/drafts/"
		}
		if _, err := os.Stat(dir + e.Longform); os.IsNotExist(err) {
			log.Infof(c, "Post no longer exists: %v", e.Longform)
			e.Delete(c)
		}
	}
}

func createPostFromLongformFile(c context.Context, dir string, file os.FileInfo, draft bool) error {
	f, err := os.Open(dir + file.Name())
	defer f.Close()
	if err != nil {
		log.Errorf(c, "Error opening file %v: %v", file.Name(), err.Error())
		return err
	}

	// get the page from file
	p, err := parser.ReadFrom(f)
	if err != nil {
		log.Errorf(c, "Error parsing file %v: %v", file.Name(), err.Error())
		return err
	}

	meta_uncast, err := p.Metadata()
	if err != nil {
		log.Errorf(c, "Error getting metadata from %v: %v", file.Name(), err.Error())
		return err
	}

	meta := map[string]string{}
	if meta_uncast != nil {
		meta, err = cast.ToStringMapStringE(meta_uncast)
		if err != nil {
			log.Errorf(c, "Error casting metadata for %v: %v. Metadata: %+v", file.Name(), err.Error(), meta_uncast)
			return err
		}
	}

	all_query := datastore.NewQuery("Entry").Filter("Longform =", file.Name()).Order("Id")
	count, err := all_query.Count(c)
	if err != nil {
		log.Warningf(c, "Error getting longform count %v: %v", file.Name(), err.Error())
	}

	if count > 1 {
		// get all from query, and delete all but the first
		entries := new([]Entry)
		_, err := all_query.GetAll(c, entries)
		if err != nil {
			log.Warningf(c, "Error getting longform duplicates %v: %v", file.Name(), err.Error())
		}

		for i, e := range *entries {
			if i >= 1 {
				e.Delete(c)
			}
		}
	}

	// create entries for those that don't exist, update those that do
	entry, err := GetLongform(c, file.Name())
	if err != nil {
		log.Warningf(c, "Error getting longform %v: %v", file.Name(), err.Error())
	}

	if entry == nil {
		entry = new(Entry)
		entry.Created = time.Now()
	}

	now := time.Now()

	year := now.Year()
	mnth := now.Month()
	day := now.Day()
	hour := now.Hour()
	min := now.Minute()

	// TODO: Don't throw away errors
	if meta["time"] != "" {
		split := strings.Split(meta["time"], ":")
		hour, _ = strconv.Atoi(split[0])
		min, _ = strconv.Atoi(split[1])
	}

	split := strings.Split(file.Name(), "-")
	year, _ = strconv.Atoi(split[0])
	m, _ := strconv.Atoi(split[1])
	mnth = time.Month(m)
	day, _ = strconv.Atoi(split[2])

	datetime := time.Date(year, mnth, day, hour, min, 0, 0, time.UTC)
	entry.Title = meta["title"]
	entry.Content = string(p.Content())
	entry.Datetime = datetime
	entry.Modified = time.Now()
	entry.Draft = draft
	entry.Longform = file.Name()

	err = entry.Save(c)
	if err != nil {
		log.Errorf(c, "Error saving entry: %v", err.Error())
		return err
	}

	return nil
}
