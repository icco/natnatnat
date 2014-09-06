package main

import (
	"appengine"
	"appengine/user"
	"github.com/pilu/traffic"
)

type NewPostPageData struct {
	Message string
}

func PostHandler(w traffic.ResponseWriter, r *traffic.Request) {
	w.Header().Set("Content-type", "text/html; charset=utf-8")
	c := appengine.NewContext(r)
	u := user.Current(c)
	if u == nil {
		url, _ := user.LoginURL(c, "/")
		fmt.Fprintf(w, `<a href="%s">Sign in or register</a>`, url)
		return
	}
	url, _ := user.LogoutURL(c, "/")
	responseData := &NewPostPageData{u}
	w.Render("new_post", responseData)
}
