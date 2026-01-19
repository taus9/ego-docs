package main

import (
	"html/template"
	"net/http"
	"path/filepath"
)

type Renderer struct {
	dir string
}

func NewRenderer(dir string) *Renderer {
	return &Renderer{dir: dir}
}

func (r *Renderer) Page(w http.ResponseWriter, name string, data any) {
	// layout + page template
	file := filepath.Join(r.dir, name)

	t, err := template.ParseFiles(file)
	if err != nil {
		http.Error(w, "template parse error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err := t.ExecuteTemplate(w, "layout", data); err != nil {
		http.Error(w, "template execute error: "+err.Error(), http.StatusInternalServerError)
		return
	}
}

func (r *Renderer) Partial(w http.ResponseWriter, name string, data any) {
	// partial only (no layout)
	file := filepath.Join(r.dir, "partials", "pages", name)

	t, err := template.ParseFiles(file)
	if err != nil {
		http.Error(w, "template parse error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err := t.Execute(w, data); err != nil {
		http.Error(w, "template execute error: "+err.Error(), http.StatusInternalServerError)
		return
	}
}
