package handlers

import (
	"appengine"
	"appengine/user"
	"errors"
	"github.com/icco/natnatnat/models"
	"github.com/pilu/traffic"
	"net/http"
	"strconv"
)

type ResponseData struct {
	Entry   *models.Entry
	IsAdmin bool
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
		responseData := &ResponseData{Entry: entry, IsAdmin: user.IsAdmin(c)}
		w.Render("post", responseData)
		return
	}
}
