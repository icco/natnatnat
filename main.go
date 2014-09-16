package main

import (
	"appengine"
	"appengine/datastore"
	"appengine/user"
	"errors"
	"github.com/gorilla/sessions"
	"github.com/pilu/traffic"
	"net/http"
	"time"
)

var router *traffic.Router
var store *sessions.CookieStore
var secret []byte

func WriteVersionKey(c appengine.Context) error {
	key := datastore.NewKey(c, "Flag", "VERSION", 0, nil)
	_, err := datastore.Put(c, key, "1.0.0")
	return err
}

func GetSessionStore(c appengine.Context) *sessions.CookieStore {

	// Unrelated action of writing Version to datastore. We do this so we can
	// look at the datastore in the AppEngine Console.
	//WriteVersionKey(c)

	default_secret := "blah blah blah, session secret for development."
	key := datastore.NewKey(c, "Flag", "SESSION_SECRET", 0, nil)
	if err := datastore.Get(c, key, &secret); err != nil {
		c.Infof("Unable to talk to Datastore: %s", err)
		secret = []byte(default_secret)
	}
	return sessions.NewCookieStore(secret)
}

func SetSessionVar(r *http.Request, w http.ResponseWriter, key string, value string) error {
	if store == nil {
		c := appengine.NewContext(r)
		store = GetSessionStore(c)
	}

	// We're ignoring the error resulted from decoding an existing session: Get()
	// always returns a session, even if empty.
	session, _ := store.Get(r, "natnatnat")
	session.Values[key] = value
	return session.Save(r, w)
}

func GetSessionVar(r *http.Request, key string) (string, error) {
	if store == nil {
		c := appengine.NewContext(r)
		store = GetSessionStore(c)
	}

	session, err := store.Get(r, "natnatnat")
	if err != nil {
		return "", err
	} else {
		data := session.Values[key]
		if str, ok := data.(string); ok {
			return str, nil
		} else {
			return "", errors.New("Not a valid key.")
		}
	}
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

// init is one of those magic functions that runs once on project create.
func init() {
	if !appengine.IsDevAppServer() {
		traffic.SetVar("env", "production")
	}

	traffic.TemplateFunc("fmttime", fmtTime)

	router = traffic.New()
	router.Get("/", RootHandler)
	router.Get("/post/new/?", NewPostGetHandler)
	router.Post("/post/new/?", NewPostPostHandler)
	router.Get("/post/:id/?", PostHandler)

	router.AddBeforeFilter(HstsMiddleware)
	router.Use(NewStaticMiddleware(traffic.PublicPath()))

	http.Handle("/", router)
}

// Entry point for go server.
func main() {}
