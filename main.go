package main

import (
	"github.com/pilu/traffic"
)

var router *traffic.Router

func init() {
	router = traffic.New()
	router.Get("/", RootHandler)
	router.Get("/post/:id/?", PostHandler)
}

func main() {
	router.Run()
}
