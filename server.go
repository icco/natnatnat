// https://golang.org/doc/articles/wiki/

package main

import (
	"errors"
	"flag"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"regexp"

	"github.com/russross/blackfriday"
)

var (
	addr = flag.Bool("addr", false, "find open address")
)

type Page struct {
	Title string
	Body  string
}

var funcMap = template.FuncMap{
	// The name "title" is what the function will be called in the template text.
	"mrkdwn": markdown,
}
var validPath = regexp.MustCompile("^/(edit|save|view)/([a-zA-Z0-9]+)$")

func markdown(text string) string {
	return string(blackfriday.MarkdownCommon([]byte(text)))
}

func (p *Page) save() error {
	filename := p.Title + ".md"
	return ioutil.WriteFile(filename, []byte(p.Body), 0600)
}

func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	f, err := ioutil.ReadFile(tmpl + ".html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	t, err := template.New(tmpl + ".html").Funcs(funcMap).Parse(string(f))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = t.Execute(w, p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func getTitle(w http.ResponseWriter, r *http.Request) (string, error) {
	m := validPath.FindStringSubmatch(r.URL.Path)
	if m == nil {
		http.NotFound(w, r)
		return "", errors.New("Invalid Page Title")
	}
	return m[2], nil // The title is the second subexpression.
}

func loadPage(title string) (*Page, error) {
	filename := title + ".md"
	text, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: string(text)}, nil
}

func makeHandler(fn func(http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Here we will extract the page title from the Request,
		// and call the provided handler 'fn'
		m := validPath.FindStringSubmatch(r.URL.Path)
		if m == nil {
			http.NotFound(w, r)
			return
		}
		fn(w, r, m[2])
	}
}

func viewHandler(w http.ResponseWriter, r *http.Request, title string) {
	p, err := loadPage(title)
	if err != nil {
		http.Redirect(w, r, "/edit/"+title, http.StatusFound)
		return
	}
	renderTemplate(w, "view", p)
}

func editHandler(w http.ResponseWriter, r *http.Request, title string) {
	p, err := loadPage(title)
	if err != nil {
		p = &Page{Title: title}
	}
	renderTemplate(w, "edit", p)
}

func saveHandler(w http.ResponseWriter, r *http.Request, title string) {
	body := r.FormValue("body")
	p := &Page{Title: title, Body: body}
	err := p.save()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/view/"+title, http.StatusFound)
}

func main() {
	flag.Parse()
	http.HandleFunc("/view/", makeHandler(viewHandler))
	http.HandleFunc("/edit/", makeHandler(editHandler))
	http.HandleFunc("/save/", makeHandler(saveHandler))

	if *addr {
		l, err := net.Listen("tcp", "127.0.0.1:0")
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("http://localhost:%s\n", []byte(l.Addr().String()))
		s := &http.Server{}
		s.Serve(l)

		return
	} else {
		fmt.Printf("http://localhost:%s\n", "8080")
		http.ListenAndServe(":8080", nil)
	}
}
