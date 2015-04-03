package handlers

import (
	"net/http"

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

type ArchiveData struct {
	DateMap map[int]Year
	IsAdmin bool
}

type Year struct {
	map[int]Month
}

type Month struct {
	map[int]Day
}

type Day struct {
	[]models.Entry
}

func ArchiveHandler(w traffic.ResponseWriter, r *traffic.Request) {
	c := appengine.NewContext(r.Request)
	entries, err := models.AllPosts(c)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	years := make(map[int]Year, 2)
	for _, p := range *entries {
		year := p.Datetime.Year()
		month := p.Datetime.Month()
		day := p.Datetime.Day()

		year_blob := years[year]
		if year_blob == nil {
			year_blob = Year{make(map[int]Month, 12)}
		}

		month_blob := year_blob[month]
		if month_blob == nil {
			month_blob = Month{make(map[int]Day, 31)}
		}

		day_blob := month_blob[day]
		if day_blob == nil {
			day_blob = Day{make([]models.Entry, 1)}
		}

		day_blob = append(day_blob, p)
	}

	data := &ArchiveData{DateMap: years, IsAdmin: user.IsAdmin(c)}
	w.Render("archive", data)
}
