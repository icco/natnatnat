package main

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/pilu/traffic"
	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
	"google.golang.org/appengine/log"
	"google.golang.org/appengine/memcache"
	"google.golang.org/appengine/taskqueue"
	"google.golang.org/appengine/user"
)

type RootData struct {
	Posts   *[]Entry
	IsAdmin bool
	Page    int64
	Prev    int64
	Next    int64
}

const perPage = 30

func RootHandler(w traffic.ResponseWriter, r *traffic.Request) {
	if r.Request.URL.Path == "/page/0" {
		http.Redirect(w, r.Request, "/", 301)
	}

	c := appengine.NewContext(r.Request)
	pg, err := strconv.ParseInt(r.Param("page"), 10, 64)
	if err != nil {
		log.Infof(c, "Error parsing: %+v", err)
		pg = 0
	}

	entries, err := Pagination(c, perPage, int(pg*perPage))
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	data := &RootData{
		Posts:   entries,
		IsAdmin: user.IsAdmin(c),
		Page:    pg,
		Next:    pg + 1,
		Prev:    pg - 1,
	}

	// If there are no posts left, don't show next button.
	if len(*entries) == 0 {
		data.Next = -1
	}

	// Get next page's posts so we don't show the next page if there is none.
	nxt_entr, err := Pagination(c, perPage, int((pg+1)*perPage))
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	if len(*nxt_entr) == 0 {
		data.Next = -1
	}

	w.Render("index", data)
}

func AboutHandler(w traffic.ResponseWriter, r *traffic.Request) {
	w.Render("about", nil)
}

func UnimplementedHandler(w traffic.ResponseWriter, r *traffic.Request) {
	http.Error(w, "Sorry, I haven't implemented this yet", 500)
}

func RedirectHomeHandler(w traffic.ResponseWriter, r *traffic.Request) {
	http.Redirect(w, r.Request, "/", 302)
}

func CleanWorkHandler(w traffic.ResponseWriter, r *traffic.Request) {
	c := appengine.NewContext(r.Request)
	q := datastore.NewQuery("Entry")
	entries := new([]Entry)
	_, err := q.GetAll(c, entries)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	for _, p := range *entries {
		// TODO: Figure out how to unset all public.
		if &p.Draft == nil {
			p.Draft = false
		}
		if len(p.Title) == 0 {
			p.Title = fmt.Sprintf("Untitled #%d", p.Id)
		}
		p.Save(c)
	}
}

type SiteMapData struct {
	Posts  *[]Entry
	Newest time.Time
}

// http://www.sitemaps.org/protocol.html
func SitemapHandler(w traffic.ResponseWriter, r *traffic.Request) {
	c := appengine.NewContext(r.Request)
	entries, err := AllPosts(c)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	data := &SiteMapData{
		Posts:  entries,
		Newest: (*entries)[0].Modified,
	}

	w.Header().Set("Content-Type", "application/xml; charset=utf-8")
	w.Render("sitemap", data)
}

// This queues lots of work every fifteen minutes.
func WorkQueueHandler(w traffic.ResponseWriter, r *traffic.Request) {
	c := appengine.NewContext(r.Request)

	// Build data for the Archive Page
	t := taskqueue.NewPOSTTask("/archive/work", url.Values{})
	_, err := taskqueue.Add(c, t, "tasks")

	if err != nil {
		log.Errorf(c, "Error queueing work: %v", err.Error())
		http.Error(w, err.Error(), 500)
		return
	}

	//// Download all the links.
	//t = taskqueue.NewPOSTTask("/link/work", url.Values{})
	//_, err = taskqueue.Add(c, t, "tasks")

	//if err != nil {
	//	log.Errorf(c, "Error queueing work: %v", err.Error())
	//	http.Error(w, err.Error(), 500)
	//	return
	//}

	// Update the stats
	t = taskqueue.NewPOSTTask("/stats/work", url.Values{})
	_, err = taskqueue.Add(c, t, "tasks")

	if err != nil {
		log.Errorf(c, "Error queueing work: %v", err.Error())
		http.Error(w, err.Error(), 500)
		return
	}

	// Update the longform data.
	t = taskqueue.NewPOSTTask("/longform/work", url.Values{})
	_, err = taskqueue.Add(c, t, "tasks")

	if err != nil {
		log.Errorf(c, "Error queueing work: %v", err.Error())
		http.Error(w, err.Error(), 500)
		return
	}

	// Clean the database
	t = taskqueue.NewPOSTTask("/clean/work", url.Values{})
	_, err = taskqueue.Add(c, t, "tasks")

	if err != nil {
		log.Errorf(c, "Error queueing work: %v", err.Error())
		http.Error(w, err.Error(), 500)
		return
	}

	// Update the Search Index
	t = taskqueue.NewPOSTTask("/search/work", url.Values{})
	_, err = taskqueue.Add(c, t, "tasks")

	if err != nil {
		log.Errorf(c, "Error queueing work: %v", err.Error())
		http.Error(w, err.Error(), 500)
		return
	}

	// Delete some caches
	hour, min, _ := time.Now().Clock()
	if hour == 0 && (min > 0 && min < 20) {
		err = memcache.Delete(c, "stats_data")
		if err != nil {
			log.Errorf(c, "Error cleaning stats cache: %v", err.Error())
		}
	}

	fmt.Fprint(w, "success.\n")
}
