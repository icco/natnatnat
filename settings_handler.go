package main

import (
	"bytes"
	"errors"
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/icco/xsrftoken"
	"github.com/pilu/traffic"

	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
	"google.golang.org/appengine/user"
)

type SettingsPageData struct {
	IsAdmin                  bool
	LogoutUrl                string
	Random                   string
	Session                  string
	TwitterAccessToken       string
	TwitterAccessTokenSecret string
	TwitterKey               string
	TwitterSecret            string
	PinboardUser             string
	PinboardToken            string
	User                     string
	Xsrf                     string
}

func SettingsGetHandler(w traffic.ResponseWriter, r *traffic.Request) {
	c := appengine.NewContext(r.Request)
	u := user.Current(c)
	if u == nil {
		url, _ := user.LoginURL(c, "/settings")
		http.Redirect(w, r.Request, url, 302)
		return
	} else {
		log.Infof(c, "Logged in as: %s", u.String())
	}

	if u != nil && !user.IsAdmin(c) {
		http.Error(w, errors.New("Not a valid user.").Error(), 403)
		return
	} else {
		url, _ := user.LogoutURL(c, "/")
		token := xsrftoken.Generate(GetFlagLogError(c, "SESSION_KEY"), u.String(), "/settings")

		twt_sec := GetFlagLogError(c, "TWITTER_SECRET")
		twt_key := GetFlagLogError(c, "TWITTER_KEY")
		twt_atok := GetFlagLogError(c, "TWITTER_ACCESS_TOKEN")
		twt_asec := GetFlagLogError(c, "TWITTER_ACCESS_SECRET")

		ses := GetFlagLogError(c, "SESSION_KEY")

		pb_usr := GetFlagLogError(c, "PINBOARD_USER")
		pb_tok := GetFlagLogError(c, "PINBOARD_TOKEN")

		responseData := &SettingsPageData{
			LogoutUrl:                url,
			Random:                   RandomString(64),
			Session:                  ses,
			PinboardUser:             pb_usr,
			PinboardToken:            pb_tok,
			TwitterAccessToken:       twt_atok,
			TwitterAccessTokenSecret: twt_asec,
			TwitterKey:               twt_key,
			TwitterSecret:            twt_sec,
			User:                     u.String(),
			Xsrf:                     token,
			IsAdmin:                  user.IsAdmin(c),
		}
		w.Render("settings", responseData)
	}
}

func SettingsPostHandler(w traffic.ResponseWriter, r *traffic.Request) {
	var err error

	c := appengine.NewContext(r.Request)
	u := user.Current(c)
	if u == nil {
		url, _ := user.LoginURL(c, "settings")
		http.Redirect(w, r.Request, url, 302)
		return
	} else {
		log.Infof(c, "Logged in as: %s", u.String())
	}

	if u != nil && !user.IsAdmin(c) {
		http.Error(w, errors.New("Not a valid user.").Error(), 403)
		return
	} else {
		xsrf := r.Request.FormValue("xsrf")

		if xsrftoken.Valid(xsrf, GetFlagLogError(c, "SESSION_KEY"), u.String(), "/settings") {
			log.Infof(c, "Valid Token!")
		} else {
			log.Infof(c, "Invalid Token...")
			http.Error(w, errors.New("Invalid Token").Error(), 403)
			return
		}

		session_key := r.Request.FormValue("session_key")
		err = SetFlag(c, "SESSION_KEY", session_key)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		twitter_key := r.Request.FormValue("twitter_key")
		err = SetFlag(c, "TWITTER_KEY", twitter_key)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		twitter_sec := r.Request.FormValue("twitter_sec")
		err = SetFlag(c, "TWITTER_SECRET", twitter_sec)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		twitter_atok := r.Request.FormValue("twitter_atok")
		err = SetFlag(c, "TWITTER_ACCESS_TOKEN", twitter_atok)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		twitter_asec := r.Request.FormValue("twitter_asec")
		err = SetFlag(c, "TWITTER_ACCESS_SECRET", twitter_asec)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		pb_usr := r.Request.FormValue("pb_usr")
		err = SetFlag(c, "PINBOARD_USER", pb_usr)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		pb_tok := r.Request.FormValue("pb_tok")
		err = SetFlag(c, "PINBOARD_TOKEN", pb_tok)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		http.Redirect(w, r.Request, "/", 302)
		return
	}
}

func RandomString(length int) string {
	buffer := bytes.NewBufferString("")

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for _, v := range r.Perm(length) {
		buffer.WriteString(fmt.Sprintf("%X", v))
	}

	return buffer.String()
}
