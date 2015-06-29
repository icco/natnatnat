package handlers

import (
	"net/http"
	"time"

	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
	"google.golang.org/appengine/user"

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
		log.Warningf(c, "Couldn't parse form: %v", r)
		http.Error(w, "Unable to parse request.", 500)
		return
	}

	in := r.Request.FormValue("text")
	md := models.Markdown(in)

	log.Infof(c, "Markdown Recieved: %s", in)
	log.Infof(c, "Markdown Rendered: %s", md)
	w.WriteText(string(md))
}

type ArchiveData struct {
	Years   map[int]Year
	Posts   *[]models.Entry
	IsAdmin bool
}

// TODO(icco): Rewrite to fix map iteration problems.
type Year map[time.Month]Month
type Month []Day
type Day []models.Entry

var months = [12]time.Month{
	time.January,
	time.February,
	time.March,
	time.April,
	time.May,
	time.June,
	time.July,
	time.August,
	time.September,
	time.October,
	time.November,
	time.December,
}

func ArchiveHandler(w traffic.ResponseWriter, r *traffic.Request) {
	c := appengine.NewContext(r.Request)
	entries, err := models.Posts(c, -1, false)
	if err != nil {
		log.Errorf(c, err.Error())
		http.Error(w, err.Error(), 500)
		return
	}
	log.Infof(c, "Retrieved data.")

	years := make(map[int]Year)

	oldest := (*entries)[0].Datetime
	newest := (*entries)[len(*entries)-1].Datetime

	for year := oldest.Year(); year <= newest.Year(); year += 1 {
		years[year] = make(Year)
		log.Infof(c, "Adding %d.", year)

		for _, month := range months {
			if year < newest.Year() || (year == newest.Year() && month <= newest.Month()) {
				years[year][month] = make([]Day, daysIn(month, year))
				log.Infof(c, "Adding %d/%d - %d days.", year, month, len(years[year][month]))
			}
		}
	}

	for _, p := range *entries {
		year := p.Datetime.Year()
		month := p.Datetime.Month()
		day := p.Datetime.Day()

		if years[year] == nil {
			log.Errorf(c, "%d isn't a valid year.", year)
		} else {
			if years[year][month] == nil {
				log.Errorf(c, "%d/%d isn't a valid month.", year, month)
			} else {
				if years[year][month][day] == nil {
					log.Errorf(c, "%d/%d/%d isn't a valid day.", year, month, day)
				} else {
					years[year][month][day] = append(years[year][month][day], p)
				}
			}
		}
	}
	log.Infof(c, "Added posts.")

	data := &ArchiveData{Years: years, IsAdmin: user.IsAdmin(c), Posts: entries}
	w.Render("archive", data)
}

// daysIn returns the number of days in a month for a given year.
func daysIn(m time.Month, year int) int {
	// This is equivalent to time.daysIn(m, year).
	return time.Date(year, m+1, 0, 0, 0, 0, 0, time.UTC).Day()
}
