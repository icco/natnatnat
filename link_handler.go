package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/pilu/traffic"
	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
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

func LinkWorkHandler(w traffic.ResponseWriter, r *traffic.Request) {
	c := appengine.NewContext(r.Request)
	user := GetFlagLogError(c, "PINBOARD_USER")
	token := GetFlagLogError(c, "PINBOARD_TOKEN")
	pb_url := fmt.Sprintf("https://api.pinboard.in/v1/%s?auth_token=%s:%s", "posts/all", user, token)

	client := urlfetch.Client(c)

	// So, there are links going back to 2003, but lets only get since 2015
	for year := time.Now().Year(); year >= 2015; year-- {
		// We've gotta sleep for five minutes so we don't get ratelimited.
		five_minutes := time.Duration(5) * time.Minute
		time.Sleep(five_minutes)

		// Ok now we do work.
		resp, err := client.Get(fmt.Sprintf("%s&fromdt=%d-01-01T00:00:00Z&todt=%d-01-01T00:00:00Z", pb_url, year, year-1))
		if err != nil {
			errorStr := "Error getting '%s': %+v. %+v"
			log.Errorf(c, errorStr, pb_url, err, resp)
			http.Error(w, fmt.Sprintf(errorStr, pb_url, err, resp), http.StatusInternalServerError)
			return
		}

		if resp.StatusCode != 200 {
			errorStr := "Error getting '%s' (status != 200): %+v"
			log.Errorf(c, errorStr, pb_url, resp.Status)
			http.Error(w, fmt.Sprintf(errorStr, pb_url, resp.Status), http.StatusInternalServerError)
			return
		}

		log.Infof(c, "Requested posts from %d.", year)
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Errorf(c, "Error reading body of '%s': %+v. '%+v' parsed from %+v", pb_url, err, body, resp)
			http.Error(w, fmt.Sprintf("Error reading body of '%s': %+v. '%+v' parsed from %+v", pb_url, err, body, resp), http.StatusInternalServerError)
			return
		}

		posts := new(PostsType)
		if err = xml.Unmarshal(body, posts); err != nil {
			log.Errorf(c, "Error parsing XML: %+v", pb_url, err)
			http.Error(w, fmt.Sprintf("Error parsing XML: %+v", pb_url, err), http.StatusInternalServerError)
			return
		}
		log.Infof(c, "Read %d posts from %d.", len(posts.Pins), year)

		for _, pin := range posts.Pins {
			tags := strings.Fields(pin.Tags)
			e := NewLink(pin.Desc, pin.Url, pin.Notes, tags, pin.Time)
			err = e.Save(c)
			if err != nil {
				log.Errorf(c, "Error saving link: %+v", pb_url, err)
				http.Error(w, fmt.Sprintf("Error saving link: %+v", pb_url, err), http.StatusInternalServerError)
				return
			}
		}
	}
}

type LinkPageData struct {
	LinkDays linkDays
	IsAdmin  bool
}

func LinkPageGetHandler(w traffic.ResponseWriter, r *traffic.Request) {
	c := appengine.NewContext(r.Request)
	lds, err := LinksByDay(c, 10)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	data := &LinkPageData{LinkDays: *lds, IsAdmin: user.IsAdmin(c)}
	w.Render("links", *data)
}
