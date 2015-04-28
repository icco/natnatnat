package handlers

import (
	"errors"
	"net/http"
	"time"

	"github.com/icco/natnatnat/models"
	"github.com/pilu/traffic"

	"code.google.com/p/xsrftoken"

	"appengine"
	"appengine/user"
)

type NewPostPageData struct {
	IsAdmin   bool
	LogoutUrl string
	User      string
	Xsrf      string
	Links     *[]models.Link
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
		token := xsrftoken.Generate(models.GetFlagLogError(c, "SESSION_KEY"), u.String(), "/post/new")
		links, err := models.Links(c, 250, true)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		responseData := &NewPostPageData{
			LogoutUrl: url,
			User:      u.String(),
			Xsrf:      token,
			IsAdmin:   user.IsAdmin(c),
			Links:     links}
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
		err := r.ParseForm()
		if err != nil {
			c.Warningf("Couldn't parse form: %v", r)
		}
		title := r.Request.FormValue("title")
		content := r.Request.FormValue("text")
		xsrf := r.Request.FormValue("xsrf")
		tags, err := models.ParseTags(content)
		public := r.Request.FormValue("private") != "on"
		if err != nil {
			c.Warningf("Couldn't parse tags: %v", err)
			tags = []string{}
		}

		c.Infof("Got POST params: title: %+v, text: %+v, xsrf: %v, private: %v", title, content, xsrf, !public)

		if xsrftoken.Valid(xsrf, models.GetFlagLogError(c, "SESSION_KEY"), u.String(), "/post/new") {
			c.Infof("Valid Token!")
		} else {
			c.Infof("Invalid Token...")
			http.Error(w, errors.New("Invalid Token").Error(), 403)
			return
		}

		e := models.NewEntry(title, content, time.Now(), public, tags)
		err = e.Save(c)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		http.Redirect(w, r.Request, e.Url(), 302)
		return
	}
}
