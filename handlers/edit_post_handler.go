package main

import (
	"errors"
	"net/http"
	"strconv"
	"time"

	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
	"google.golang.org/appengine/user"

	"github.com/icco/xsrftoken"
	"github.com/pilu/traffic"
)

type EditPostPageData struct {
	Entry     *Entry
	IsAdmin   bool
	LogoutUrl string
	User      string
	Xsrf      string
	EditUrl   string
}

func EditPostGetHandler(w traffic.ResponseWriter, r *traffic.Request) {
	c := appengine.NewContext(r.Request)
	id, err := strconv.ParseInt(r.Param("id"), 10, 64)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	entry, err := GetEntry(c, id)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	u := user.Current(c)
	if u == nil {
		url, _ := user.LoginURL(c, entry.EditUrl())
		http.Redirect(w, r.Request, url, 302)
		return
	} else {
		log.Infof(c, "Logged in as: %s", u.String())
	}

	if u != nil && !user.IsAdmin(c) {
		http.Error(w, errors.New("Not a valid user.").Error(), 403)
		return
	} else {
		url, _ := user.LogoutURL(c, "/")
		token := xsrftoken.Generate(GetFlagLogError(c, "SESSION_KEY"), u.String(), entry.EditUrl())
		responseData := &EditPostPageData{
			LogoutUrl: url,
			User:      u.String(),
			Xsrf:      token,
			IsAdmin:   user.IsAdmin(c),
			EditUrl:   entry.EditUrl(),
			Entry:     entry}
		w.Render("edit_post", responseData)
	}
}

func EditPostPostHandler(w traffic.ResponseWriter, r *traffic.Request) {
	c := appengine.NewContext(r.Request)
	id, err := strconv.ParseInt(r.Param("id"), 10, 64)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	entry, err := GetEntry(c, id)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	u := user.Current(c)
	if u == nil {
		url, _ := user.LoginURL(c, entry.EditUrl())
		http.Redirect(w, r.Request, url, 302)
		return
	} else {
		log.Infof(c, "Logged in as: %s", u.String())
	}

	if u != nil && !user.IsAdmin(c) {
		http.Error(w, errors.New("Not a valid user.").Error(), 403)
		return
	} else {
		err := r.ParseForm()
		if err != nil {
			log.Warningf(c, "Couldn't parse form: %v", r)
		}
		title := r.Request.FormValue("title")
		content := r.Request.FormValue("text")
		xsrf := r.Request.FormValue("xsrf")
		tags, err := ParseTags(content)
		public := r.Request.FormValue("private") != "on"
		if err != nil {
			log.Warningf(c, "Couldn't parse tags: %v", err)
			tags = []string{}
		}

		log.Infof(c, "Got POST params: title: %+v, text: %+v, xsrf: %v, private: %v", title, content, xsrf, !public)

		if xsrftoken.Valid(xsrf, GetFlagLogError(c, "SESSION_KEY"), u.String(), entry.EditUrl()) {
			log.Infof(c, "Valid Token!")
		} else {
			log.Infof(c, "Invalid Token...")
			http.Error(w, errors.New("Invalid Token").Error(), 403)
			return
		}

		entry.Modified = time.Now()
		entry.Title = title
		entry.Tags = tags
		entry.Content = content
		entry.Public = public

		err = entry.Save(c)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		http.Redirect(w, r.Request, entry.Url(), 302)
		return
	}
}
