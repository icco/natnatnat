package handlers

import (
	"net/http"
	"time"

	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
	"google.golang.org/appengine/log"
	"google.golang.org/appengine/user"

	"github.com/icco/natnatnat/models"
	"github.com/pilu/traffic"
)

type ArchiveData struct {
	Years   map[int]Year
	Posts   *[]models.Entry
	IsAdmin bool
}

// TODO(icco): Rewrite to fix map iteration problems.
type Year map[time.Month]Month
type Month []Day
type Day []int64

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

	entries, err := models.AllPosts(c)
	if err != nil {
		log.Errorf(c, err.Error())
		http.Error(w, err.Error(), 500)
		return
	}
	log.Infof(c, "Retrieved data: %d.", len(*entries))

	years := make(map[int]Year)

	oldest := (*entries)[len(*entries)-1].Datetime
	newest := (*entries)[0].Datetime

	log.Infof(c, "Oldest: %v, Newest: %v", oldest, newest)

	// for year := oldest.Year(); year <= newest.Year(); year += 1 {
	// 	years[year] = make(Year)
	// 	log.Infof(c, "Adding %d.", year)
	// 	for _, month := range months {
	// 		if year < newest.Year() || (year == newest.Year() && month <= newest.Month()) {
	// 			years[year][month] = make([]Day, daysIn(month, year))
	// 			log.Debugf(c, "Adding %d/%d - %d days.", year, month, len(years[year][month]))
	// 		}
	// 	}
	// }

	q := models.ArchivePageQuery()
	t := q.Run(c)
	for {
		var p models.Entry
		_, err := t.Next(&p)
		if err == datastore.Done {
			break // No further entities match the query.
		}

		if err != nil {
			log.Errorf(c, "Error fetching next Entry: %v", err)
			break
		}

		year := p.Datetime.Year()
		month := p.Datetime.Month()
		day := p.Datetime.Day()
		log.Infof(c, "Trying post id %d", p.Id)

		if years[year] == nil {
			years[year] = make(Year)
			log.Errorf(c, "%d isn't a valid year.", year)
		}

		if years[year][month] == nil {
			log.Errorf(c, "%d/%d isn't a valid month.", year, month)
			years[year][month] = make([]Day, daysIn(month, year))
		}

		if years[year][month][day] == nil {
			log.Infof(c, "Making %d/%d/%d", year, month, day)
			years[year][month][day] = make(Day, 0)
		}

		log.Infof(c, "Appending %d/%d/%d: %+v", year, month, day, years[year][month][day])
		years[year][month][day] = append(years[year][month][day], p.Id)
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
