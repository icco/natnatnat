package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"sort"
	"strings"
	"time"

	"github.com/pilu/traffic"

	"google.golang.org/appengine"
	"google.golang.org/appengine/taskqueue"
	"google.golang.org/appengine/urlfetch"
	"google.golang.org/appengine/user"
)

/*
$ curl https://user:passwd@api.pinboard.in/v1/posts/recent

    <?xml version="1.0" encoding="UTF-8" ?>
    <posts dt="2011-03-25T14:49:56Z" user="user">
        <post href="http://www.slate.com/" description="Slate"
        extended="online news and comment"  hash="3c56b6c6cfedbe75f41e79e6fa102aba"
        tag="news opinion" time="2011-03-24T20:30:47Z" />
        ...
    </posts>
*/
type LinkXML struct {
	XMLName xml.Name  `xml:"post"`
	Url     string    `xml:"href,attr"`
	Desc    string    `xml:"description,attr"`
	Notes   string    `xml:"extended,attr"`
	Time    time.Time `xml:"time,attr"`
	Hash    string    `xml:"hash,attr"`
	Shared  bool      `xml:"shared,attr"`
	Tags    string    `xml:"tag,attr"`
	Meta    string    `xml:"meta,attr"`
}

type PostsType struct {
	XMLName xml.Name  `xml:"posts"`
	Pins    []LinkXML `xml:"post"`
}

func LinkQueueHandler(w traffic.ResponseWriter, r *traffic.Request) {
	c := appengine.NewContext(r.Request)
	t := taskqueue.NewPOSTTask("/link/work", url.Values{})
	_, err := taskqueue.Add(c, t, "")

	if err != nil {
		http.Error(w, err.Error(), 500)
	} else {
		fmt.Fprint(w, "success.\n")
	}
}

func LinkWorkHandler(w traffic.ResponseWriter, r *traffic.Request) {
	c := appengine.NewContext(r.Request)
	user := GetFlagLogError(c, "PINBOARD_USER")
	token := GetFlagLogError(c, "PINBOARD_TOKEN")
	params := "count=100"
	pb_url := fmt.Sprintf("https://api.pinboard.in/v1/%s?auth_token=%s:%s&%s", "posts/recent", user, token, params)

	client := urlfetch.Client(c)
	resp, err := client.Get(pb_url)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error getting '%s': %+v", pb_url, err), http.StatusInternalServerError)
		return
	}

	if resp.StatusCode != 200 {
		http.Error(w, fmt.Sprintf("Error getting '%s': %+v", pb_url, resp.Status), http.StatusInternalServerError)
		return
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error reading body of '%s': %+v. '%+v' parsed from %+v", pb_url, err, body, resp), http.StatusInternalServerError)
		return
	}

	posts := new(PostsType)
	if err = xml.Unmarshal(body, posts); err != nil {
		http.Error(w, fmt.Sprintf("Error parsing XML: %+v", pb_url, err), http.StatusInternalServerError)
		return
	}

	for _, pin := range posts.Pins {
		tags := strings.Fields(pin.Tags)
		e := NewLink(pin.Desc, pin.Url, pin.Notes, tags, pin.Time)
		err = e.Save(c)
		if err != nil {
			http.Error(w, fmt.Sprintf("Error saving link: %+v", pb_url, err), http.StatusInternalServerError)
			return
		}
	}
}

type LinkPageData struct {
	LinkDays linkDays
	IsAdmin  bool
}

type LinkDay struct {
	Links []Link
	Day   time.Time
}

type linkDays []*LinkDay

// These three functions are needed for Sort.
func (p linkDays) Len() int           { return len(p) }
func (p linkDays) Less(i, j int) bool { return p[i].Day.Before(p[j].Day) }
func (p linkDays) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p linkDays) HasDate(day time.Time) int {
	for i, d := range p {
		if d.Day == day {
			return i
		}
	}
	return -1
}

func LinkPageGetHandler(w traffic.ResponseWriter, r *traffic.Request) {
	c := appengine.NewContext(r.Request)
	links, err := AllLinks(c)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	days := make([]*LinkDay, 0)

	for _, l := range *links {
		date := l.Posted.Round(time.Hour * 24)
		if linkDays(days).HasDate(date) < 0 {
			days = append(days, &LinkDay{
				Day:   date,
				Links: []Link{l},
			})
		} else {
			i := linkDays(days).HasDate(date)
			days[i].Links = append(days[i].Links, l)
		}
	}

	lds := linkDays(days)
	sort.Reverse(lds)

	data := &LinkPageData{LinkDays: lds, IsAdmin: user.IsAdmin(c)}
	w.Render("links", data)
}
