package main

import (
	"html/template"
	"net/http"
	"time"

	"google.golang.org/appengine"
	"google.golang.org/appengine/user"

	"golang.org/x/net/context"

	"github.com/gorilla/sessions"
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
	return Markdown(args...)
}

func monthToInt(m string) int {
	for _, mnth := range Months {
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
	router.Get("/", RootHandler)
	router.Get("/page/:page/?", RootHandler)

	router.Get("/about", AboutHandler)
	router.Get("/stats", StatsHandler)
	router.Get("/posts.json", StatsHistoryJsonHandler)
	router.Get("/sitemap.xml", SitemapHandler)

	router.Get("/archive(s?)", ArchiveHandler)
	router.Get("/archive/queue", ArchiveQueueHandler)
	router.Post("/archive/work", ArchiveTaskHandler)

	router.Post("/md", MarkdownHandler)

	router.Get("/admin/?", AdminGetHandler)

	router.Get("/post/new/?", NewPostGetHandler)
	router.Post("/post/new/?", NewPostPostHandler)

	router.Get("/post/:id/?", PostHandler)

	router.Get("/edit/:id/?", EditPostGetHandler)
	router.Post("/edit/:id/?", EditPostPostHandler)

	router.Get("/tags/:id/?", TagHandler)
	router.Get("/tags/?", TagsHandler)

	router.Get("/day/:year/:month/:day/?", DayHandler)

	router.Get("/aliases", TagAliasGetHandler)
	router.Post("/aliases", TagAliasPostHandler)

	router.Get("/settings", SettingsGetHandler)
	router.Post("/settings", SettingsPostHandler)

	router.Get("/mention", WebMentionGetHandler)
	router.Post("/mention", WebMentionPostHandler)

	router.Get("/feed.atom", FeedAtomHandler)
	router.Get("/feed.rss", FeedRssHandler)

	router.Get("/summary.atom", SummaryAtomHandler)
	router.Get("/summary.rss", SummaryRssHandler)

	router.Get("/link/queue", LinkQueueHandler)
	router.Post("/link/work", LinkWorkHandler)
	router.Get("/links", LinkPageGetHandler)

	// TODO: REMOVE.
	router.Get("/import", ImportPseudowebHandler)

	router.AddBeforeFilter(HstsMiddleware)
	router.Use(NewStaticMiddleware(traffic.PublicPath()))

	http.Handle("/", router)
}

// Entry point for go server.
func main() {}
