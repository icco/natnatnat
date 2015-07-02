package main

import (
	"html/template"
	"net/http"
	"time"

	"google.golang.org/appengine"
	"google.golang.org/appengine/user"

	"golang.org/x/net/context"

	"github.com/gorilla/sessions"
	"github.com/icco/natnatnat/handlers"
	"github.com/icco/natnatnat/models"
	"github.com/pilu/traffic"
)

var store *sessions.CookieStore

func HstsMiddleware(w traffic.ResponseWriter, r *traffic.Request) {
	w.Header().Add("Strict-Transport-Security", "max-age=15768000")
}

func IsAdmin(c context.Context) bool {
	return c != nil && user.IsAdmin(c)
}

func fmtTime(t time.Time) string {
	const layout = "03:04 on Jan 2, 2006 UTC"
	return t.Format(layout)
}

func jsonTime(t time.Time) string {
	b, err := t.MarshalText()
	if err != nil {
		// TODO(icco): Log something
		return ""
	}

	return string(b)
}

func markdown(args ...interface{}) template.HTML {
	return models.Markdown(args...)
}

func monthToInt(m string) int {
	for _, mnth := range handlers.Months {
		if mnth.String() == m {
			return int(mnth)
		}
	}

	return -1
}

// init is one of those magic functions that runs once on project create.
func init() {
	if !appengine.IsDevAppServer() {
		traffic.SetVar("env", "production")
	}

	traffic.TemplateFunc("fmttime", fmtTime)
	traffic.TemplateFunc("jsontime", jsonTime)
	traffic.TemplateFunc("mrkdwn", markdown)
	traffic.TemplateFunc("m2i", monthToInt)

	router := traffic.New()
	router.Get("/", handlers.RootHandler)
	router.Get("/page/:page/?", handlers.RootHandler)

	router.Get("/about", handlers.AboutHandler)
	router.Get("/stats", handlers.StatsHandler)
	router.Get("/posts.json", handlers.StatsHistoryJsonHandler)

	router.Get("/archive(s?)", handlers.ArchiveHandler)
	router.Get("/archive/queue", handlers.ArchiveQueueHandler)
	router.Post("/archive/work", handlers.ArchiveTaskHandler)

	router.Post("/md", handlers.MarkdownHandler)

	router.Get("/post/new/?", handlers.NewPostGetHandler)
	router.Post("/post/new/?", handlers.NewPostPostHandler)

	router.Get("/post/:id/?", handlers.PostHandler)

	router.Get("/edit/:id/?", handlers.EditPostGetHandler)
	router.Post("/edit/:id/?", handlers.EditPostPostHandler)

	router.Get("/tags/:id/?", handlers.TagHandler)
	router.Get("/tags/?", handlers.TagsHandler)

	router.Get("/day/:year/:month/:day/?", handlers.DayHandler)

	router.Get("/aliases", handlers.TagAliasGetHandler)
	router.Post("/aliases", handlers.TagAliasPostHandler)

	router.Get("/settings", handlers.SettingsGetHandler)
	router.Post("/settings", handlers.SettingsPostHandler)

	router.Get("/mention", handlers.WebMentionGetHandler)
	router.Post("/mention", handlers.WebMentionPostHandler)

	router.Get("/feed.atom", handlers.FeedAtomHandler)
	router.Get("/feed.rss", handlers.FeedRssHandler)

	router.Get("/summary.atom", handlers.SummaryAtomHandler)
	router.Get("/summary.rss", handlers.SummaryRssHandler)

	router.Get("/link/queue", handlers.LinkQueueHandler)
	router.Post("/link/work", handlers.LinkWorkHandler)
	router.Get("/links", handlers.LinkPageGetHandler)

	router.AddBeforeFilter(HstsMiddleware)
	router.Use(NewStaticMiddleware(traffic.PublicPath()))

	http.Handle("/", router)
}

// Entry point for go server.
func main() {}
