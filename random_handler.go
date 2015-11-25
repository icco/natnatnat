package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
	"google.golang.org/appengine/user"

	"github.com/pilu/traffic"
)

type ImportStruct struct {
	Datetime time.Time `json:"date"`
	Text     string    `json:"text"`
	Title    string    `json:"title"`
	Link     string    `josn:"link"`
}

func ImportPseudowebHandler(w traffic.ResponseWriter, r *traffic.Request) {
	c := appengine.NewContext(r.Request)
	u := user.Current(c)
	if u == nil {
		url, _ := user.LoginURL(c, "/")
		http.Redirect(w, r.Request, url, 302)
		return
	} else {
		log.Infof(c, "Logged in as: %s", u.String())
	}

	if u != nil && !user.IsAdmin(c) {
		http.Error(w, errors.New("Not a valid user.").Error(), 403)
		return
	} else {
		file, err := ioutil.ReadFile("tmp/all.json")
		if err != nil {
			e := fmt.Sprintf("File error: %v", err)
			log.Errorf(c, e)
			http.Error(w, e, 500)
			return
		}

		var data []ImportStruct
		json.Unmarshal(file, &data)
		log.Debugf(c, "Loaded: %v", data)

		for _, p := range data {
			p.Text = fmt.Sprintf("Posted originally at %s.\n\n%s", p.Link, p.Text)
			e := NewEntry(p.Title, p.Text, p.Datetime, true, []string{})
			e.Save(c)
		}

		w.WriteText("Finished.")
	}
}
