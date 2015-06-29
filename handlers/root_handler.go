package handlers

import (
	"net/http"

	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
	"google.golang.org/appengine/user"

	"github.com/icco/natnatnat/models"
	"github.com/pilu/traffic"
)

type RootData struct {
	Posts   interface{}
	IsAdmin bool
}

func RootHandler(w traffic.ResponseWriter, r *traffic.Request) {
	c := appengine.NewContext(r.Request)
	entries, err := models.AllPosts(c)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	data := &RootData{Posts: entries, IsAdmin: user.IsAdmin(c)}
	w.Render("index", data)
}

func AboutHandler(w traffic.ResponseWriter, r *traffic.Request) {
	http.Redirect(w, r.Request, "http://natwelch.com", 301)
}

func UnimplementedHandler(w traffic.ResponseWriter, r *traffic.Request) {
	http.Error(w, "Sorry, I haven't implemented this yet", 500)
}

func MarkdownHandler(w traffic.ResponseWriter, r *traffic.Request) {
	c := appengine.NewContext(r.Request)

	err := r.ParseForm()
	if err != nil {
		log.Warningf(c, "Couldn't parse form: %v", r)
		http.Error(w, "Unable to parse request.", 500)
		return
	}

	in := r.Request.FormValue("text")
	md := models.Markdown(in)

	log.Infof(c, "Markdown Recieved: %s", in)
	log.Infof(c, "Markdown Rendered: %s", md)
	w.WriteText(string(md))
}
