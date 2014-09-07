package main

import (
	"appengine"
	"appengine/user"
	"code.google.com/p/xsrftoken"
	"fmt"
	"github.com/pilu/traffic"
	"net/http"
	"time"
)

type NewPostPageData struct {
	LogoutUrl string
	User      string
	Xsrf      string
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

	// key is a secret key for your application. userID is a unique identifier
	// for the user. actionID is the action the user is taking (e.g. POSTing to a
	// particular path).
	token := xsrftoken.Generate(string(secret), u.String(), "/post/new")
	SetSessionVar(r.Request, w, "xsrf", token)
	responseData := &NewPostPageData{LogoutUrl: url, User: u.String(), Xsrf: token}
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

	title := r.Request.FormValue("title")
	content := r.Request.FormValue("text")
	xsrf := r.Request.FormValue("xsrf")

	c.Infof("Got POST params: title: %+v, text: %+v, xsrf: %v", title, content, xsrf)
	if xsrftoken.Valid(xsrf, string(secret), u.String(), "/post/new") {
		c.Infof("Valid Token!")
	} else {
		w.WriteHeader(403)
		return
	}

	e := NewEntry(title, content, time.Now(), []string{})
	e.save(c)

	new_route := fmt.Sprintf("/post/%d", e.Id)

	http.Redirect(w, r.Request, new_route, 302)
	return
}
