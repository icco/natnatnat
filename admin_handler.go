package main

import (
	"errors"
	"net/http"

	"github.com/pilu/traffic"

	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
	"google.golang.org/appengine/user"
)

type AdminPageData struct {
	IsAdmin   bool
	LogoutUrl string
	User      string
	Drafts    *[]Entry
	Longform  *[]Entry
}

func AdminGetHandler(w traffic.ResponseWriter, r *traffic.Request) {
	c := appengine.NewContext(r.Request)
	u := user.Current(c)
	if u == nil {
		url, _ := user.LoginURL(c, "/admin")
		http.Redirect(w, r.Request, url, 302)
		return
	} else {
		log.Infof(c, "Logged in as: %s", u.String())
	}

	if u != nil && !user.IsAdmin(c) {
		http.Error(w, errors.New("Not a valid user.").Error(), 403)
		return
	} else {
		url, err := user.LogoutURL(c, "/")
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		drafts, err := Drafts(c)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		lngfrm, err := LongformPosts(c)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		responseData := &AdminPageData{
			LogoutUrl: url,
			User:      u.String(),
			IsAdmin:   user.IsAdmin(c),
			Drafts:    drafts,
			Longform:  lngfrm,
		}
		w.Render("admin", responseData)
	}
}
