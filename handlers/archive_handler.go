package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
	"google.golang.org/appengine/memcache"
	"google.golang.org/appengine/taskqueue"
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
type Day int64

var Months = [12]time.Month{
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

func ArchiveQueueHandler(w traffic.ResponseWriter, r *traffic.Request) {
	c := appengine.NewContext(r.Request)
	t := taskqueue.NewPOSTTask("/archive/work", url.Values{})
	_, err := taskqueue.Add(c, t, "")

	if err != nil {
		http.Error(w, err.Error(), 500)
	} else {
		fmt.Fprint(w, "success.\n")
	}
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
		for _, month := range Months {
			notFirstOrLastYear := year < newest.Year() || year > oldest.Year()
			notMonthsInFuture := year != newest.Year() || month <= newest.Month()
			notMonthsBeforeOldest := year != oldest.Year() || month >= oldest.Month()
			if notFirstOrLastYear || (notMonthsInFuture && notMonthsBeforeOldest) {
				mstr := month.String()
				years[ystr][mstr] = make([]Day, daysIn(month, year)+1)
				log.Debugf(c, "Adding %d/%d - %d days.", year, month, len(years[ystr][mstr]))

				for day := range years[ystr][mstr] {
					if day > 0 {
						e, err := models.PostsForDay(c, int64(year), int64(month), int64(day))
						if err != nil {
							log.Errorf(c, err.Error())
							http.Error(w, err.Error(), 500)
							return
						}
						years[ystr][mstr][day] = Day(len(*e))
					}
				}
			}
		}
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
