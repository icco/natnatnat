package handlers

import (
	"net/http"
	"time"

	"appengine"
	"appengine/user"

	"github.com/icco/natnatnat/models"
	"github.com/pilu/traffic"
)

type RootData struct {
	Posts   interface{}
	IsAdmin bool
}

func RootHandler(w traffic.ResponseWriter, r *traffic.Request) {
	c := appengine.NewContext(r.Request)
	entries, err := models.AllPosts(c)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	data := &RootData{Posts: entries, IsAdmin: user.IsAdmin(c)}
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
		c.Warningf("Couldn't parse form: %v", r)
		http.Error(w, "Unable to parse request.", 500)
		return
	}

	in := r.Request.FormValue("text")
	md := models.Markdown(in)

	c.Infof("Markdown Recieved: %s", in)
	c.Infof("Markdown Rendered: %s", md)
	w.WriteText(string(md))
}

type ArchiveData struct {
	Years   map[int]Year
	IsAdmin bool
}

// TODO(icco): Rewrite to fix map iteration problems.
type Year map[time.Month]Month
type Month map[int]Day
type Day []models.Entry

func ArchiveHandler(w traffic.ResponseWriter, r *traffic.Request) {
	c := appengine.NewContext(r.Request)
	entries, err := models.Posts(c, -1, false)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	years := make(map[int]Year)

	for _, p := range *entries {
		year := p.Datetime.Year()
		month := p.Datetime.Month()
		day := p.Datetime.Day()

		if years[year] == nil {
			years[year] = make(Year)
		}

		if years[year][month] == nil {
			years[year][month] = make(Month)
		}

		if years[year][month][day] == nil {
			years[year][month][day] = make(Day, 0)
		}

		years[year][month][day] = append(years[year][month][day], p)
	}

	data := &ArchiveData{Years: years, IsAdmin: user.IsAdmin(c)}
	w.Render("archive", data)
}
