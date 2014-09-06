package main

import (
	"github.com/pilu/traffic"
)

type RootData struct {
	Message string
}

func RootHandler(w traffic.ResponseWriter, r *traffic.Request) {
	responseData := &RootData{"Hello World"}
	w.Render("index", responseData)
}
