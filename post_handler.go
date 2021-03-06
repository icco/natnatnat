package main

import (
	"bytes"
	"errors"
	"net/http"
	"strconv"
	"text/template"
	"time"

	"github.com/pilu/traffic"
	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
	"google.golang.org/appengine/user"
)

type ResponseData struct {
	Entry   *Entry
	IsAdmin bool
	Next    string
	Prev    string
}

func PostHandler(w traffic.ResponseWriter, r *traffic.Request) {
	c := appengine.NewContext(r.Request)
	id, err := strconv.ParseInt(r.Param("id"), 10, 64)
	if err != nil {
		log.Errorf(c, "Error parsing int: %+v", err)
		http.Redirect(w, r.Request, "/", 302)
		return
	}
	entry, err := GetEntry(c, id)
	if err != nil {
		log.Errorf(c, "Error loading post: %+v", err)
		http.Error(w, err.Error(), 500)
		return
	}

	if entry.Draft && !user.IsAdmin(c) {
		http.Error(w, errors.New("Post is not public").Error(), 403)
		return
	} else {
		responseData := &ResponseData{
			Entry:   entry,
			IsAdmin: user.IsAdmin(c),
			Next:    entry.NextPost(c),
			Prev:    entry.PrevPost(c)}
		w.Render("post", responseData)
		return
	}
}

type DayData struct {
	Posts   *[]Entry
	IsAdmin bool
	Date    time.Time
}

func DayHandler(w traffic.ResponseWriter, r *traffic.Request) {
	c := appengine.NewContext(r.Request)
	day, err := strconv.ParseInt(r.Param("day"), 10, 32)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	month, err := strconv.ParseInt(r.Param("month"), 10, 32)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	year, err := strconv.ParseInt(r.Param("year"), 10, 32)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	entries, err := PostsForDay(c, year, month, day)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	responseData := &DayData{
		Posts:   entries,
		IsAdmin: user.IsAdmin(c),
		Date:    time.Date(int(year), time.Month(month), int(day), 0, 0, 0, 0, time.UTC),
	}
	w.Render("day", responseData)
	return
}

func PostMarkdownHandler(w traffic.ResponseWriter, r *traffic.Request) {
	c := appengine.NewContext(r.Request)
	id, err := strconv.ParseInt(r.Param("id"), 10, 64)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	entry, err := GetEntry(c, id)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	if entry.Draft && !user.IsAdmin(c) {
		http.Error(w, errors.New("Post is not public").Error(), 403)
		return
	}

	// Define a template.
	const md = `---

id: {{.Id}}
datetime: "{{.Datetime}}"
title: "{{if .Title}}{{.Title}}{{else}}#{{.Id}}{{end}}"
draft: {{.Draft}}
permalink: "/post/{{.Id}}"
{{if .Longform}}longform: {{.Longform}}
{{end}}
---

{{.Content}}
`

	// Create a new template and parse the letter into it.
	t := template.Must(template.New("post_md").Parse(md))

	buf := new(bytes.Buffer)
	err = t.Execute(buf, entry)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.WriteText(buf.String())
	return
}

func PostJsonHandler(w traffic.ResponseWriter, r *traffic.Request) {
	c := appengine.NewContext(r.Request)
	id, err := strconv.ParseInt(r.Param("id"), 10, 64)
	if err != nil {
		log.Errorf(c, "Error loading post: %+v", err)
		http.Error(w, err.Error(), 500)
		return
	}
	entry, err := GetEntry(c, id)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	if entry.Draft && !user.IsAdmin(c) {
		http.Error(w, errors.New("Post is not public").Error(), 403)
		return
	}

	post := &EntryJson{
		Html:     entry.Html(),
		ReadTime: entry.ReadTime(),
	}
	post.Id = entry.Id
	post.Title = entry.Title
	post.Content = entry.Content
	post.Datetime = entry.Datetime
	post.Created = entry.Created
	post.Modified = entry.Modified
	post.Tags = entry.Tags

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteJSON(post)
}
