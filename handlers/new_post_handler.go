package handlers

import (
	"appengine"
	"appengine/user"
	"code.google.com/p/xsrftoken"
	"errors"
	"fmt"
	"github.com/icco/natnatnat/models"
	"github.com/pilu/traffic"
	"net/http"
	"strings"
	"time"
)

type NewPostPageData struct {
	IsAdmin   bool
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

	if u != nil && !user.IsAdmin(c) {
		http.Error(w, errors.New("Not a valid user.").Error(), 403)
		return
	} else {
		url, _ := user.LogoutURL(c, "/")
		token := xsrftoken.Generate(models.GetFlag(c, "SESSION_KEY"), u.String(), "/post/new")
		responseData := &NewPostPageData{LogoutUrl: url, User: u.String(), Xsrf: token, IsAdmin: user.IsAdmin(c)}
		w.Render("new_post", responseData)
	}
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

	if u != nil && !user.IsAdmin(c) {
		http.Error(w, errors.New("Not a valid user.").Error(), 403)
		return
	} else {
		title := r.Request.FormValue("title")
		content := r.Request.FormValue("text")
		xsrf := r.Request.FormValue("xsrf")
		tags := strings.Split(r.Request.FormValue("tags"), ",")

		// TODO(icco): Add as HTML option
		public := true

		c.Infof("Got POST params: title: %+v, text: %+v, xsrf: %v", title, content, xsrf)
		if xsrftoken.Valid(xsrf, models.GetFlag(c, "SESSION_KEY"), u.String(), "/post/new") {
			c.Infof("Valid Token!")
		} else {
			c.Infof("Invalid Token...")
			http.Error(w, errors.New("Invalid Token").Error(), 403)
			return
		}

		e := models.NewEntry(title, content, time.Now(), public, tags)
		err := e.Save(c)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		new_route := fmt.Sprintf("/post/%d", e.Id)
		http.Redirect(w, r.Request, new_route, 302)
		return
	}
}
