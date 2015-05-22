package handlers

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"appengine"
	"appengine/user"

	"code.google.com/p/xsrftoken"

	"github.com/icco/natnatnat/models"
	"github.com/pilu/traffic"
)

type TagData struct {
	Posts interface{}
	Tag   string
}

func TagHandler(w traffic.ResponseWriter, r *traffic.Request) {
	c := appengine.NewContext(r.Request)
	tag := r.Param("id")

	if tag == "" {
		http.Redirect(w, r.Request, "/tags", 301)
	}

	if tag != strings.ToLower(tag) {
		http.Redirect(w, r.Request, fmt.Sprintf("/tags/%s", strings.ToLower(tag)), 301)
	}

	entries, err := models.PostsWithTag(c, tag)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	data := &TagData{Posts: entries, Tag: tag}
	w.Render("tag", data)
}

type TagsData struct {
	Tags map[string]int
}

func TagsHandler(w traffic.ResponseWriter, r *traffic.Request) {
	c := appengine.NewContext(r.Request)
	w.Render("tags", &TagsData{Tags: models.AllTags(c)})
}

func TagAliasGetHandler(w traffic.ResponseWriter, r *traffic.Request) {
	c := appengine.NewContext(r.Request)
	u := user.Current(c)
	if u == nil {
		url, _ := user.LoginURL(c, "/post/new")
		http.Redirect(w, r.Request, url, 302)
		return
	} else {
		c.Infof("Logged in as: %s", u.String())
	}

	if u != nil && !user.IsAdmin(c) {
		http.Error(w, errors.New("Not a valid user.").Error(), 403)
		return
	} else {
		w.Render("aliases", &TagsData{Tags: models.AllTags(c)})
		return
	}
}

func TagAliasPostHandler(w traffic.ResponseWriter, r *traffic.Request) {
	c := appengine.NewContext(r.Request)
	u := user.Current(c)
	if u == nil {
		url, _ := user.LoginURL(c, "/post/new")
		http.Redirect(w, r.Request, url, 302)
		return
	} else {
		c.Infof("Logged in as: %s", u.String())
	}

	if u != nil && !user.IsAdmin(c) {
		http.Error(w, errors.New("Not a valid user.").Error(), 403)
		return
	} else {
		err := r.ParseForm()
		if err != nil {
			c.Warningf("Couldn't parse form: %v", r)
		}
		xsrf := r.Request.FormValue("xsrf")

		if xsrftoken.Valid(xsrf, models.GetFlagLogError(c, "SESSION_KEY"), u.String(), r.Request.URL.Path) {
			c.Infof("Valid Token!")
		} else {
			c.Infof("Invalid Token...")
			http.Error(w, errors.New("Invalid Token").Error(), 403)
			return
		}

		http.Redirect(w, r.Request, "/aliases", 302)
		return
	}
}
