package main

import (
	"appengine"
	"appengine/user"
	"github.com/pilu/traffic"
	"net/http"
)

type NewPostPageData struct {
	LogoutUrl string
	User      string
}

func NewPostGetHandler(w traffic.ResponseWriter, r *traffic.Request) {
	c := appengine.NewContext(r.Request)
	u := user.Current(c)
	if u == nil {
		url, _ := user.LoginURL(c, "/post/new")
		http.Redirect(w, r.Request, url, 302)
		return
	} else {
		c.Infof("Logged in as: %s", u.String())
	}
	url, _ := user.LogoutURL(c, "/")
	responseData := &NewPostPageData{LogoutUrl: url, User: u.String()}
	w.Render("new_post", responseData)
}

func NewPostPostHandler(w traffic.ResponseWriter, r *traffic.Request) {
	c := appengine.NewContext(r.Request)
	u := user.Current(c)
	if u == nil {
		url, _ := user.LoginURL(c, "/post/new")
		http.Redirect(w, r.Request, url, 302)
		return
	} else {
		c.Infof("Logged in as: %s", u.String())
	}
	c.Infof("Got POST params: title: %+v, text: %+v", r.Request.FormValue("title"), r.Request.FormValue("text"))
	http.Redirect(w, r.Request, "/", 302)
	return
}
