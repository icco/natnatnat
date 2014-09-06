package main

import (
	"appengine"
	"appengine/user"
	"fmt"
	"github.com/pilu/traffic"
)

type NewPostPageData struct {
	LogoutUrl string
	User      string
}

func NewPostHandler(w traffic.ResponseWriter, r *traffic.Request) {
	w.Header().Set("Content-type", "text/html; charset=utf-8")
	c := appengine.NewContext(r.Request)
	u := user.Current(c)
	if u == nil {
		url, _ := user.LoginURL(c, "/")
		fmt.Fprintf(w, `<a href="%s">Sign in or register</a>`, url)
		return
	}
	url, _ := user.LogoutURL(c, "/")
	responseData := &NewPostPageData{LogoutUrl: url, User: u.String()}
	w.Render("new_post", responseData)
}
