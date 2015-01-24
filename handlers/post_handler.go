package handlers

import (
	"appengine"
	"appengine/user"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/icco/natnatnat/models"
	"github.com/pilu/traffic"
)

type ResponseData struct {
	Entry   *models.Entry
	IsAdmin bool
	Next    string
	Prev    string
}

func PostHandler(w traffic.ResponseWriter, r *traffic.Request) {
	c := appengine.NewContext(r.Request)
	id, err := strconv.ParseInt(r.Param("id"), 10, 64)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	entry, err := models.GetEntry(c, id)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	if !entry.Public {
		http.Error(w, errors.New("Post is not public").Error(), 403)
		return
	} else {
		responseData := &ResponseData{
			Entry:   entry,
			IsAdmin: user.IsAdmin(c),
			Next:    entry.NextPost(c),
			Prev:    entry.PrevPost(c)}
		w.Render("post", responseData)
		return
	}
}
