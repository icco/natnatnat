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
	Posts   *[]Entry
	Tag     string
	Aliases []string
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

	isAlias, alias := models.GetAlias(c, tag)
	if isAlias {
		http.Redirect(w, r.Request, fmt.Sprintf("/tags/%s", alias), 301)
	}

	entries, err := models.PostsWithTag(c, tag)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	aliases := models.GetTagAliases(c, tag)
	data := &TagData{Posts: entries, Tag: tag, Aliases: *aliases}
	w.Render("tag", data)
}

type TagsData struct {
	Tags map[string]int
}

func TagsHandler(w traffic.ResponseWriter, r *traffic.Request) {
	c := appengine.NewContext(r.Request)
	w.Render("tags", &TagsData{Tags: models.AllTags(c)})
}

type AliasData struct {
	Aliases map[string]string
	Xsrf    string
	IsAdmin bool
}

func TagAliasGetHandler(w traffic.ResponseWriter, r *traffic.Request) {
	c := appengine.NewContext(r.Request)
	u := user.Current(c)
	if u == nil {
		url, _ := user.LoginURL(c, "/aliases")
		http.Redirect(w, r.Request, url, 302)
		return
	} else {
		c.Infof("Logged in as: %s", u.String())
	}

	if u != nil && !user.IsAdmin(c) {
		http.Error(w, errors.New("Not a valid user.").Error(), 403)
		return
	} else {
		token := xsrftoken.Generate(models.GetFlagLogError(c, "SESSION_KEY"), u.String(), "/aliases")
		w.Render("aliases", &AliasData{
			Aliases: models.AliasMap(c),
			Xsrf:    token,
			IsAdmin: user.IsAdmin(c),
		})
		return
	}
}

func TagAliasPostHandler(w traffic.ResponseWriter, r *traffic.Request) {
	c := appengine.NewContext(r.Request)
	u := user.Current(c)
	if u == nil {
		url, _ := user.LoginURL(c, "/aliases")
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
		from := r.Request.FormValue("name")
		to := r.Request.FormValue("tag")

		if xsrftoken.Valid(xsrf, models.GetFlagLogError(c, "SESSION_KEY"), u.String(), r.Request.URL.Path) {
			c.Infof("Valid Token!")
			a := models.NewAlias(from, to)
			err = a.Save(c)
			if err != nil {
				http.Error(w, err.Error(), 500)
				return
			}
		} else {
			c.Infof("Invalid Token...")
			http.Error(w, errors.New("Invalid Token").Error(), 403)
			return
		}

		http.Redirect(w, r.Request, "/aliases", 302)
		return
	}
}
