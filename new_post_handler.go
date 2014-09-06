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
	w.Header().Set("Content-type", "text/html; charset=utf-8")
	c := appengine.NewContext(r.Request)
	u := user.Current(c)
	if u == nil {
		url, _ := user.LoginURL(c, "/")
		http.Redirect(w, r.Request, url, 302)
		return
	}
	url, _ := user.LogoutURL(c, "/")
	responseData := &NewPostPageData{LogoutUrl: url, User: u.String()}
	w.Render("new_post", responseData)
}

func NewPostPostHandler(w traffic.ResponseWriter, r *traffic.Request) {
	w.Header().Set("Content-type", "text/html; charset=utf-8")
	c := appengine.NewContext(r.Request)
	u := user.Current(c)
	if u == nil {
		url, _ := user.LoginURL(c, "/")
		http.Redirect(w, r.Request, url, 302)
		return
	}
	url, _ := user.LogoutURL(c, "/")
	responseData := &NewPostPageData{LogoutUrl: url, User: u.String()}
	w.Render("new_post", responseData)
}
