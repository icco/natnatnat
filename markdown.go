package main

import (
	"html/template"
	"regexp"

	"github.com/russross/blackfriday"
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

func twitterHandleToMarkdown(in []byte) []byte {
	return TwitterHandleRegex.ReplaceAll(in, []byte("$1[@$2](http://twitter.com/$2)"))
}

func hashTagsToMarkdown(in []byte) []byte {
	return HashtagRegex.ReplaceAll(in, []byte("$1[#$2](/tags/$2)"))
}
