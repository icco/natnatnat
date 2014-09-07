package main

import (
	"appengine"
	"github.com/pilu/traffic"
	"net/http"
	"strconv"
)

type ResponseData struct {
	Entry *Entry
}

func PostHandler(w traffic.ResponseWriter, r *traffic.Request) {
	c := appengine.NewContext(r.Request)
	id, err := strconv.ParseInt(r.Param("id"), 10, 64)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	entry, err := GetEntry(c, id)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	responseData := &ResponseData{entry}
	w.Render("post", responseData)
}
