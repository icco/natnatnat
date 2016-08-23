package main

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/pilu/traffic"
	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
	"google.golang.org/appengine/log"
	"google.golang.org/appengine/memcache"
	"google.golang.org/appengine/user"
)

type StatsData struct {
	Posts        int
	WordsPerPost float64
	PostsPerDay  float64
	WordsPerDay  float64
	DaysSince    float64
	IsAdmin      bool
	LinksPerPost float64
	LinksPerDay  float64
	Years        []string
	YearData     map[string][]float64
}

func StatsHandler(w traffic.ResponseWriter, r *traffic.Request) {
	c := appengine.NewContext(r.Request)

	var data *StatsData
	if json_data, err := memcache.Get(c, "stats-data"); err == memcache.ErrCacheMiss {
		entries, err := AllPosts(c)
		if err != nil {
			log.Errorf(c, "Error loading posts: %+v", err)
			http.Error(w, err.Error(), 500)
			return
		}

		postCount := len(*entries)
		oldestPost := (*entries)[postCount-1]
		dayCount := time.Since(oldestPost.Datetime).Hours() / 24.0

		years := []string{}
		yearData := make(map[string][]float64)
		for i := time.Now().Year(); i >= oldestPost.Datetime.Year(); i-- {
			years = append(years, strconv.Itoa(i))
			yearData[strconv.Itoa(i)] = []float64{0.0, 0.0, 0.0}
		}

		words := 0
		for _, p := range *entries {
			words += len(strings.Fields(p.Content))
			words += len(strings.Fields(p.Title))
			yearData[strconv.Itoa(p.Datetime.Year())][0] += 1
		}

		for _, y := range years {
			yearData[y][1] = yearData[y][0] / 52.0
		}

		readLinks := 0
		q := LinkQuery(c, -1, true)
		for t := q.Run(c); ; {
			var l Link
			_, err := t.Next(&l)
			if err == datastore.Done {
				break
			}
			if err != nil {
				log.Errorf(c, "Error loading link: %+v", err)
			} else {
				readLinks += 1
				yearData[strconv.Itoa(l.Posted.Year())][2] += 1
			}
		}

		data = &StatsData{
			Posts:        postCount,
			PostsPerDay:  float64(postCount) / dayCount,
			WordsPerPost: float64(words) / float64(postCount),
			WordsPerDay:  float64(words) / dayCount,
			DaysSince:    dayCount,
			LinksPerDay:  float64(readLinks) / dayCount,
			Years:        years,
			YearData:     yearData,
		}

		b, err := json.Marshal(data)
		if err != nil {
			log.Errorf(c, err.Error())
			http.Error(w, err.Error(), 500)
			return
		}
		item := &memcache.Item{
			Key:   "stats-data",
			Value: b,
		}

		// Set the item, unconditionally
		if err := memcache.Set(c, item); err != nil {
			log.Errorf(c, "error setting item: %v", err)
		}
	} else if err != nil {
		log.Errorf(c, "error getting item: %v", err)
	} else {
		err := json.Unmarshal(json_data.Value, &data)
		if err != nil {
			log.Errorf(c, err.Error())
			http.Error(w, err.Error(), 500)
			return
		}
	}

	data.IsAdmin = user.IsAdmin(c)
	w.Render("stats", data)
}

func StatsHistoryJsonHandler(w traffic.ResponseWriter, r *traffic.Request) {
	c := appengine.NewContext(r.Request)
	entries, err := AllPosts(c)
	if err != nil {
		log.Errorf(c, "Error loading posts: %+v", err)
		http.Error(w, err.Error(), 500)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteJSON(entries)
}
