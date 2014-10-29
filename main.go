package main

import (
	"appengine"
	"appengine/user"
	"github.com/gorilla/sessions"
	"github.com/icco/natnatnat/handlers"
	"github.com/icco/natnatnat/models"
	"github.com/pilu/traffic"
	"html/template"
	"net/http"
	"time"
)

var store *sessions.CookieStore

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
	return models.Markdown(args...)
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
	router.Get("/about", handlers.AboutHandler)

	router.Get("/post/new/?", handlers.NewPostGetHandler)
	router.Post("/post/new/?", handlers.NewPostPostHandler)

	router.Get("/post/:id/?", handlers.PostHandler)

	router.Get("/edit/:id/?", handlers.EditPostGetHandler)
	router.Post("/edit/:id/?", handlers.EditPostPostHandler)

	router.Get("/settings", handlers.SettingsGetHandler)
	router.Post("/settings", handlers.SettingsPostHandler)

	router.Get("/mention", handlers.WebMentionGetHandler)
	router.Post("/mention", handlers.WebMentionPostHandler)

	router.Get("/feed.atom", handlers.FeedAtomHandler)
	router.Get("/feed.rss", handlers.FeedRssHandler)

	router.AddBeforeFilter(HstsMiddleware)
	router.Use(NewStaticMiddleware(traffic.PublicPath()))

	http.Handle("/", router)
}

// Entry point for go server.
func main() {}
