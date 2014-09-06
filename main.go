package main

import (
	"github.com/pilu/traffic"
	"net/http"
)

var router *traffic.Router

// init is one of those magic functions that runs once on project create.
func init() {
	router = traffic.New()
	router.Get("/", RootHandler)
	router.Get("/post/new/?", NewPostGetHandler)
	router.Post("/post/new/?", NewPostPostHandler)
	router.Get("/post/:id/?", PostHandler)
	http.Handle("/", router)
}

// Entry point for go server.
func main() {}
