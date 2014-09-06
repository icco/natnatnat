package main

import (
	"github.com/pilu/traffic"
)

var router *traffic.Router

// init is one of those magic functions that runs once on project create.
func init() {
	router = traffic.New()
	router.Get("/", RootHandler)
	router.Get("/post/new/?", NewPostHandler)
	router.Get("/post/:id/?", PostHandler)
}

// Entry point for go server.
func main() {
	router.Run()
}
