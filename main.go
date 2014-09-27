package main

import (
	"appengine"
	"appengine/user"
	"bytes"
	"fmt"
	"github.com/gorilla/sessions"
	"github.com/icco/natnatnat/handlers"
	"github.com/pilu/traffic"
	"github.com/russross/blackfriday"
	"html/template"
	"math/rand"
	"net/http"
	"time"
)

var store *sessions.CookieStore
var secret []byte

func RandomString(length int) string {
	buffer := bytes.NewBufferString("")

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for _, v := range r.Perm(length) {
		buffer.WriteString(fmt.Sprintf("%X", v))
	}

	return buffer.String()
}

func WriteVersionKey(c appengine.Context) error {
	return SetFlag(c, "VERSION", "1.0.1")
}

func HstsMiddleware(w traffic.ResponseWriter, r *traffic.Request) {
	w.Header().Add("Strict-Transport-Security", "max-age=15768000")
}

func IsAdmin(c appengine.Context) bool {
	return c != nil && user.IsAdmin(c)
}

func fmtTime(t time.Time) string {
	const layout = "03:04 on Jan 2, 2006"
	return t.Format(layout)
}

func markdown(args ...interface{}) template.HTML {
	inc := []byte(fmt.Sprintf("%s", args...))
	s := blackfriday.MarkdownCommon(inc)
	return template.HTML(s)
}

// init is one of those magic functions that runs once on project create.
func init() {
	if !appengine.IsDevAppServer() {
		traffic.SetVar("env", "production")
	}

	traffic.TemplateFunc("fmttime", fmtTime)
	traffic.TemplateFunc("mrkdwn", markdown)

	router := traffic.New()
	router.Get("/", handlers.RootHandler)

	router.Get("/post/new/?", NewPostGetHandler)
	router.Post("/post/new/?", NewPostPostHandler)

	router.Get("/post/:id/?", PostHandler)

	router.Get("/settings", SettingsGetHandler)
	router.Post("/settings", SettingsPostHandler)

	router.AddBeforeFilter(HstsMiddleware)
	router.Use(NewStaticMiddleware(traffic.PublicPath()))

	http.Handle("/", router)
}

// Entry point for go server.
func main() {}
