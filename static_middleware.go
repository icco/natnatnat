package main

import (
	"github.com/pilu/traffic"
	"log"
	"mime"
	"net/http"
	"path/filepath"
)

type StaticMiddleware struct {
	publicPath string
}

func (middleware *StaticMiddleware) ServeHTTP(w traffic.ResponseWriter, r *traffic.Request, next traffic.NextMiddlewareFunc) {
	callNext := func() {
		if nextMiddleware := next(); nextMiddleware != nil {
			nextMiddleware.ServeHTTP(w, r, next)
		}
	}

	dir := http.Dir(middleware.publicPath)
	path := r.URL.Path

	file, err := dir.Open(path)
	if err != nil {
		callNext()
		return
	}
	defer file.Close()

	if info, err := file.Stat(); err == nil && !info.IsDir() {
		ctype := mime.TypeByExtension(filepath.Ext(path))
		log.Printf("content-type according to go: %s", ctype)
		w.Header().Set("Content-Type", ctype)
		log.Printf("Static file headers: %s", w.Header())
		http.ServeContent(w, r.Request, path, info.ModTime(), file)
		return
	}

	callNext()
}

func NewStaticMiddleware(publicPath string) *StaticMiddleware {
	middleware := &StaticMiddleware{
		publicPath: publicPath,
	}

	return middleware
}
