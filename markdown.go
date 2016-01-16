package main

import (
	"html/template"
	"net/http"
	"regexp"
	"strings"

	"github.com/pilu/traffic"
	"github.com/russross/blackfriday"
	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
)

var HashtagRegex *regexp.Regexp = regexp.MustCompile(`(\s)#(\w+)`)
var TwitterHandleRegex *regexp.Regexp = regexp.MustCompile(`(\s)@([_A-Za-z0-9]+)`)

// Markdown generator.
func Markdown(str string) template.HTML {
	inc := []byte(str)
	inc = twitterHandleToMarkdown(inc)
	inc = hashTagsToMarkdown(inc)
	s := blackfriday.MarkdownCommon(inc)
	return template.HTML(s)
}

// Takes a chunk of markdown and just returns the first paragraph.
func SummarizeText(str string) string {
	out := strings.Split(str, "\n")
	return out[0]
}

func twitterHandleToMarkdown(in []byte) []byte {
	return TwitterHandleRegex.ReplaceAll(in, []byte("$1[@$2](http://twitter.com/$2)"))
}

func hashTagsToMarkdown(in []byte) []byte {
	return HashtagRegex.ReplaceAll(in, []byte("$1[#$2](/tags/$2)"))
}

type MarkdownHandlerData struct {
	Text template.HTML
}

func MarkdownHandler(w traffic.ResponseWriter, r *traffic.Request) {
	c := appengine.NewContext(r.Request)

	err := r.ParseForm()
	if err != nil {
		log.Warningf(c, "Couldn't parse form: %v", r)
		http.Error(w, "Unable to parse request.", 500)
		return
	}

	in := r.Request.FormValue("text")
	md := Markdown(in)

	log.Infof(c, "Markdown Recieved: %s", in)
	log.Infof(c, "Markdown Rendered: %s", md)
	w.Render("blank", &MarkdownHandlerData{Text: md})
}
