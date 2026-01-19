package main

import "net/http"

type App struct {
	Renderer *Renderer
}

func (a *App) Home(w http.ResponseWriter, r *http.Request) {
	// render the full layout page
	a.Renderer.Page(w, "layout.html", nil) // if your layout is the homepage itself
}

func (a *App) GettingStarted(w http.ResponseWriter, r *http.Request) {
	a.Renderer.Partial(w, "start.html", nil)
}

func (a *App) LanguageBasics(w http.ResponseWriter, r *http.Request) {
	a.Renderer.Partial(w, "basics.html", nil)
}

func (a *App) EgoErrors(w http.ResponseWriter, r *http.Request) {
	a.Renderer.Partial(w, "errors.html", nil)
}

func (a *App) Functions(w http.ResponseWriter, r *http.Request) {
	a.Renderer.Partial(w, "functions.html", nil)
}

func (a *App) Structures(w http.ResponseWriter, r *http.Request) {
	a.Renderer.Partial(w, "structures.html", nil)
}

func (a *App) ScopeAndLifetime(w http.ResponseWriter, r *http.Request) {
	a.Renderer.Partial(w, "scope.html", nil)
}

func (a *App) ControlFlow(w http.ResponseWriter, r *http.Request) {
	a.Renderer.Partial(w, "control.html", nil)
}

func (a *App) ExpressionsAndOperators(w http.ResponseWriter, r *http.Request) {
	a.Renderer.Partial(w, "expressions.html", nil)
}

func (a *App) StandardLibrary(w http.ResponseWriter, r *http.Request) {
	a.Renderer.Partial(w, "stdlib.html", nil)
}

func (a *App) VariablesAndConstants(w http.ResponseWriter, r *http.Request) {
	a.Renderer.Partial(w, "variables.html", nil)
}

func (a *App) ValuesAndTypes(w http.ResponseWriter, r *http.Request) {
	a.Renderer.Partial(w, "values.html", nil)
}
