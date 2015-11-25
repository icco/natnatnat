package main

import (
	"net/http"
	"strconv"
	"time"

	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
	"google.golang.org/appengine/user"

	"github.com/pilu/traffic"
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

	if len(*entries) == 0 {
		data.Next = -1
	}

	w.Render("index", data)
}

func AboutHandler(w traffic.ResponseWriter, r *traffic.Request) {
	http.Redirect(w, r.Request, "http://natwelch.com", 301)
}

func UnimplementedHandler(w traffic.ResponseWriter, r *traffic.Request) {
	http.Error(w, "Sorry, I haven't implemented this yet", 500)
}

func MarkdownHandler(w traffic.ResponseWriter, r *traffic.Request) {
	c := appengine.NewContext(r.Request)

	err := r.ParseForm()
	if err != nil {
		log.Warningf(c, "Couldn't parse form: %v", r)
		http.Error(w, "Unable to parse request.", 500)
		return
	}

	in := r.Request.FormValue("text")
	md := Markdown(in)

	log.Infof(c, "Markdown Recieved: %s", in)
	log.Infof(c, "Markdown Rendered: %s", md)
	w.WriteText(string(md))
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
