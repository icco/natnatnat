package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
	"google.golang.org/appengine/log"
	"google.golang.org/appengine/memcache"
	"google.golang.org/appengine/user"

	"github.com/icco/natnatnat/models"
	"github.com/pilu/traffic"
)

type ArchiveData struct {
	Years   *map[string]Year
	Months  *[]string
	Posts   *[]models.Entry
	IsAdmin bool
}

type Year map[string]Month
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

func ArchiveTaskHandler(w traffic.ResponseWriter, r *traffic.Request) {
	c := appengine.NewContext(r.Request)

	entries, err := models.AllPosts(c)
	if err != nil {
		log.Errorf(c, err.Error())
		http.Error(w, err.Error(), 500)
		return
	}
	log.Infof(c, "Retrieved data: %d.", len(*entries))

	years := make(map[string]Year)

	oldest := (*entries)[len(*entries)-1].Datetime
	newest := (*entries)[0].Datetime

	log.Infof(c, "Oldest: %v, Newest: %v", oldest, newest)

	for year := oldest.Year(); year <= newest.Year(); year += 1 {
		ystr := strconv.Itoa(year)
		years[ystr] = make(Year)
		log.Infof(c, "Adding %d.", year)
		for _, month := range months {
			if year < newest.Year() || (year == newest.Year() && month <= newest.Month()) {
				mstr := month.String()
				years[ystr][mstr] = make([]Day, daysIn(month, year))
				log.Debugf(c, "Adding %d/%d - %d days.", year, month, len(years[ystr][mstr]))
			}
		}
	}

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

		year := strconv.Itoa(p.Datetime.Year())
		yint := p.Datetime.Year()
		month := p.Datetime.Month()
		mstr := month.String()
		day := p.Datetime.Day()
		log.Infof(c, "Trying post id %d", p.Id)

		if years[year] == nil {
			years[year] = make(Year)
			log.Errorf(c, "%s isn't a valid year.", year)
		}

		if years[year][mstr] == nil {
			log.Errorf(c, "%s/%d isn't a valid month.", year, month)
			years[year][mstr] = make([]Day, daysIn(month, yint))
		}

		if years[year][mstr][day] == nil {
			log.Infof(c, "Making %s/%d/%d", year, month, day)
			years[year][mstr][day] = make(Day, 0)
		}

		years[year][mstr][day] = append(years[year][mstr][day], p.Id)
	}
	log.Infof(c, "Added posts.")

	// https://blog.golang.org/json-and-go
	b, err := json.Marshal(years)
	if err != nil {
		log.Errorf(c, err.Error())
		http.Error(w, err.Error(), 500)
	}

	item := &memcache.Item{
		Key:   "archive_data",
		Value: b,
	}

	// Set the item, unconditionally
	if err := memcache.Set(c, item); err != nil {
		log.Errorf(c, "error setting item: %v", err)
	}
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

	// Get the item from the memcache
	var years map[string]Year
	if year_data, err := memcache.Get(c, "archive_data"); err == memcache.ErrCacheMiss {
		log.Infof(c, "item not in the cache")
	} else if err != nil {
		log.Errorf(c, "error getting item: %v", err)
	} else {
		err := json.Unmarshal(year_data.Value, &years)
		if err != nil {
			log.Errorf(c, err.Error())
			http.Error(w, err.Error(), 500)
			return
		}
	}

	data := &ArchiveData{
		Years: &years,
		Months: &[]string{
			time.January.String(),
			time.February.String(),
			time.March.String(),
			time.April.String(),
			time.May.String(),
			time.June.String(),
			time.July.String(),
			time.August.String(),
			time.September.String(),
			time.October.String(),
			time.November.String(),
			time.December.String(),
		},
		IsAdmin: user.IsAdmin(c),
		Posts:   entries,
	}
	w.Render("archive", data)
}

// daysIn returns the number of days in a month for a given year.
func daysIn(m time.Month, year int) int {
	// This is equivalent to time.daysIn(m, year).
	return time.Date(year, m+1, 0, 0, 0, 0, 0, time.UTC).Day()
}
