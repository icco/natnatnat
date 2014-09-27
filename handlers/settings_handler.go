package handlers

import (
	"appengine"
	"appengine/user"
	"code.google.com/p/xsrftoken"
	"errors"
	"github.com/icco/natnatnat/models"
	"github.com/pilu/traffic"
	"net/http"
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
	User                     string
	Version                  string
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
		token := xsrftoken.Generate(string(secret), u.String(), "/settings")

		twt_sec, _ := models.GetFlag(c, "TWITTER_SECRET")
		twt_key, _ := models.GetFlag(c, "TWITTER_KEY")
		twt_atok, _ := models.GetFlag(c, "TWITTER_ACCESS_TOKEN")
		twt_asec, _ := models.GetFlag(c, "TWITTER_ACCESS_SECRET")

		ses, _ := models.GetFlag(c, "SESSION_KEY")
		ver, _ := models.GetFlag(c, "VERSION")
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		responseData := &SettingsPageData{
			LogoutUrl:                url,
			Random:                   RandomString(64),
			Session:                  ses,
			TwitterAccessToken:       twt_atok,
			TwitterAccessTokenSecret: twt_asec,
			TwitterKey:               twt_key,
			TwitterSecret:            twt_sec,
			User:                     u.String(),
			Version:                  ver,
			Xsrf:                     token,
			IsAdmin:                  IsAdmin(c),
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
		c.Infof("Logged in as: %s", u.String())
	}

	if u != nil && !user.IsAdmin(c) {
		http.Error(w, errors.New("Not a valid user.").Error(), 403)
		return
	} else {
		xsrf := r.Request.FormValue("xsrf")

		if xsrftoken.Valid(xsrf, string(secret), u.String(), "/settings") {
			c.Infof("Valid Token!")
		} else {
			c.Infof("Invalid Token...")
			http.Error(w, errors.New("Invalid Token").Error(), 403)
			return
		}

		session_key := r.Request.FormValue("session_key")
		err = models.SetFlag(c, "SESSION_KEY", session_key)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		twitter_key := r.Request.FormValue("twitter_key")
		err = models.SetFlag(c, "TWITTER_KEY", twitter_key)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		twitter_sec := r.Request.FormValue("twitter_sec")
		err = models.SetFlag(c, "TWITTER_SECRET", twitter_sec)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		twitter_atok := r.Request.FormValue("twitter_atok")
		err = models.SetFlag(c, "TWITTER_ACCESS_TOKEN", twitter_atok)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		twitter_asec := r.Request.FormValue("twitter_asec")
		err = models.SetFlag(c, "TWITTER_ACCESS_SECRET", twitter_asec)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		http.Redirect(w, r.Request, "/", 302)
		return
	}
}
