package main

import (
	"log"
	"net/http"
)

func main() {
	r := NewRenderer("./web/templates")
	app := &App{Renderer: r}

	mux := http.NewServeMux()

	// full page
	mux.HandleFunc("/", app.Home)

	// partials for HTMX
	mux.HandleFunc("/docs/start", app.GettingStarted)
	mux.HandleFunc("/docs/basics", app.LanguageBasics)
	mux.HandleFunc("/docs/errors", app.EgoErrors)
	mux.HandleFunc("/docs/functions", app.Functions)
	mux.HandleFunc("/docs/structures", app.Structures)
	mux.HandleFunc("/docs/scope", app.ScopeAndLifetime)
	mux.HandleFunc("/docs/control", app.ControlFlow)
	mux.HandleFunc("/docs/expressions", app.ExpressionsAndOperators)
	mux.HandleFunc("/docs/stdlib", app.StandardLibrary)
	mux.HandleFunc("/docs/variables", app.VariablesAndConstants)
	mux.HandleFunc("/docs/values", app.ValuesAndTypes)

	// static files
	fileServer := http.FileServer(http.Dir("./web/static/"))
	mux.Handle("/static/", http.StripPrefix("/static/", fileServer))

	log.Println("listening on :4000")
	log.Fatal(http.ListenAndServe(":4000", mux))
}
