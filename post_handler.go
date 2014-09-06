package main

import (
	"github.com/pilu/traffic"
)

type ResponseData struct {
	Message string
}

func PostHandler(w traffic.ResponseWriter, r *traffic.Request) {
	responseData := &ResponseData{r.Param("id")}
	w.Render("index", responseData)
}
