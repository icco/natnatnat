package main

import (
	"appengine"
	"appengine/user"
	"fmt"
	"github.com/gorilla/sessions"
	"github.com/icco/natnatnat/handlers"
	"github.com/pilu/traffic"
	"github.com/russross/blackfriday"
	"html/template"
	"net/http"
	"regexp"
	"time"
)

var store *sessions.CookieStore
var TwitterHandleRegex *regexp.Regexp = regexp.MustCompile(`(?<=^|(?<=[^a-zA-Z0-9-_\.]))@([A-Za-z]+[A-Za-z0-9]+)`)

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
	inc = twitterHandleToMarkdown(inc)
	s := blackfriday.MarkdownCommon(inc)
	return template.HTML(s)
}

func twitterHandleToMarkdown(in []byte) []byte {
	return TwitterHandleRegex.ReplaceAll(in, []byte("[@$2](http://twitter.com/$2)"))
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

	router.Get("/post/new/?", handlers.NewPostGetHandler)
	router.Post("/post/new/?", handlers.NewPostPostHandler)

	router.Get("/post/:id/?", handlers.PostHandler)

	router.Get("/settings", handlers.SettingsGetHandler)
	router.Post("/settings", handlers.SettingsPostHandler)

	router.Get("/mention", handlers.WebMentionGetHandler)
	router.Post("/mention", handlers.WebMentionPostHandler)

	router.AddBeforeFilter(HstsMiddleware)
	router.Use(NewStaticMiddleware(traffic.PublicPath()))

	http.Handle("/", router)
}

// Entry point for go server.
func main() {}
