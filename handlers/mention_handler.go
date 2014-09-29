package handlers

/*
User Aaron posts a blog post on his blog
User Barnaby writes post on his blog that links to Aaron's post.
After publishing the post (i.e. it has a URL), Barnaby's server notices this link as part of the publishing process
Barnaby's server does webmention discovery on Aaron's post to find its webmention endpoint (if not found, process stops)
Barnaby's server sends a webmention to Aaron's post's webmention endpoint with
source set to Barnaby's post's permalink
target set to Aaron's post's permalink.
Aaron's server receives the webmention
Aaron's server verifies that target (after following redirects) in the webmention is a valid permalink on Aaron's blog (if not, processing stops)
Aaron's server verifies that the source (when retrieved, after following redirects) in the webmention contains a hyperlink to the target (if not, processing stops)
Unmentioned but implied (and vaguely mentioned in the pingback spec):
Aaron's server displays the information about Barnaby's post somewhere on Aaron's post.
*/

import (
	"appengine"
	"github.com/icco/natnatnat/models"
	"github.com/pilu/traffic"
)

func WebMentionGetHandler(w traffic.ResponseWriter, r *traffic.Request) {
	c := appengine.NewContext(r.Request)
	err := models.WriteVersionKey(c)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}

func WebMentionPostHandler(w traffic.ResponseWriter, r *traffic.Request) {
	c := appengine.NewContext(r.Request)
	err := models.WriteVersionKey(c)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}
