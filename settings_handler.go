package main

import (
	"appengine"
	"appengine/user"
	"code.google.com/p/xsrftoken"
	"errors"
	"fmt"
	"github.com/pilu/traffic"
	"net/http"
	"strings"
	"time"
)

type SettingsPageData struct {
	LogoutUrl string
	User      string
	Xsrf      string
}

func SettingsGetHandler(w traffic.ResponseWriter, r *traffic.Request) {
	c := appengine.NewContext(r.Request)
	u := user.Current(c)
	if u == nil {
		url, _ := user.LoginURL(c, "/settings")
		http.Redirect(w, r.Request, url, 302)
		return
	} else {
		c.Infof("Logged in as: %s", u.String())
	}

	if u != nil && !user.IsAdmin(c) {
		http.Error(w, errors.New("Not a valid user.").Error(), 403)
		return
	} else {
		err := WriteVersionKey(c)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		url, _ := user.LogoutURL(c, "/")
		// key is a secret key for your application. userID is a unique identifier
		// for the user. actionID is the action the user is taking (e.g. POSTing to a
		// particular path).
		token := xsrftoken.Generate(string(secret), u.String(), "/settings")
		responseData := &SettingsPageData{LogoutUrl: url, User: u.String(), Xsrf: token}
		w.Render("settings", responseData)
	}
}

func SettingsPostHandler(w traffic.ResponseWriter, r *traffic.Request) {
	c := appengine.NewContext(r.Request)
	u := user.Current(c)
	if u == nil {
		url, _ := user.LoginURL(c, "settings")
		http.Redirect(w, r.Request, url, 302)
		return
	} else {
		c.Infof("Logged in as: %s", u.String())
	}

	if u != nil && !user.IsAdmin(c) {
		http.Error(w, errors.New("Not a valid user.").Error(), 403)
		return
	} else {
		session_key := r.Request.FormValue("session_key")
		twitter_key := r.Request.FormValue("twitter_key")
		xsrf := r.Request.FormValue("xsrf")

		if xsrftoken.Valid(xsrf, string(secret), u.String(), "/settings") {
			c.Infof("Valid Token!")
		} else {
			c.Infof("Invalid Token...")
			http.Error(w, errors.New("Invalid Token").Error(), 403)
			return
		}

		err := WriteVersionKey(c)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		http.Redirect(w, r.Request, "/", 302)
		return
	}
}
