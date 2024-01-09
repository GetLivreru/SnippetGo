package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/justinas/alice"
	"snippetbox.KaiyrbekovAdilet.net/ui"
)

func (app *application) routes() http.Handler {
	router := httprouter.New()
	
	router.NotFound = http.HandlerFunc(func (w http.ResponseWriter, r *http.Request)  {
		app.notFound(w)
	})	
	
	fileServer := http.FileServer(http.FS(ui.Files))
	router.Handler(http.MethodGet, "/static/*filepath", fileServer)
	router.HandlerFunc(http.MethodGet, "/ping", ping)

	dynamic := alice.New(app.sessionManager.LoadAndSave, noSurf, app.authenticate)

	 
	protected := dynamic.Append(app.requireAuthentication)

	router.Handler(http.MethodGet, "/snippet/create", protected.ThenFunc(app.snippetCreate))
	router.Handler(http.MethodPost, "/snippet/create", protected.ThenFunc(app.snippetCreatePost))
	router.Handler(http.MethodGet, "/account/view",protected.ThenFunc(app.accountView))
	router.Handler(http.MethodPost, "/user/logout", dynamic.ThenFunc(app.userLogoutPost))
	router.Handler(http.MethodPost, "/account/password/update", dynamic.ThenFunc(app.accountPasswordUpdatePost))
	router.Handler(http.MethodGet, "/account/password/update", dynamic.ThenFunc(app.accountPasswordUpdate))

	standard := alice.New(app.recoverPanic, app.logRequest, secureHeaders)
	return standard.Then(router)
}