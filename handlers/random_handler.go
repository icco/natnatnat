package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
	"google.golang.org/appengine/user"

	"github.com/icco/natnatnat/models"
	"github.com/pilu/traffic"
)

type ImportStruct struct {
	Id       int       `json:"id"`
	Datetime time.Time `json:"datetime"`
	Text     string    `json:"text"`
	Title    string    `json:"title"`
}

func ImportTumbleHandler(w traffic.ResponseWriter, r *traffic.Request) {
	c := appengine.NewContext(r.Request)
	u := user.Current(c)
	md := models.Markdown("test")
	if u == nil {
		url, _ := user.LoginURL(c, "/post/new")
		http.Redirect(w, r.Request, url, 302)
		return
	} else {
		log.Infof(c, "Logged in as: %s", u.String())
	}

	if u != nil && !user.IsAdmin(c) {
		http.Error(w, errors.New("Not a valid user.").Error(), 403)
		return
	} else {
		file, err := ioutil.ReadFile("tumbledata.json")
		if err != nil {
			e := fmt.Sprintf("File error: %v", err)
			http.Error(w, e, 500)
			return
		}

		var data []ImportStruct
		json.Unmarshal(file, &data)
		log.Debugf("Loaded: %v", data)

		w.WriteText("Finished.")
	}
}
