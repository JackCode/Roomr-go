package main

import (
	"net/http"

	"github.com/bmizerany/pat"
	"github.com/justinas/alice"
)

func (app *application) routes() http.Handler {
	standardMiddleware := alice.New(app.recoverPanic, app.logRequest, secureHeaders)
	dynamicMiddleware := alice.New(app.session.Enable, noSurf)

	mux := pat.New()
	mux.Get("/", dynamicMiddleware.ThenFunc(app.home))

	return standardMiddleware.Then(mux)
}
