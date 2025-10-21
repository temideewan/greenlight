package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) routes() http.Handler {
	// initialize a new http router instance
	router := httprouter.New()

	// use custom not found handler for the httprouter notfound
	router.NotFound = http.HandlerFunc(app.notFoundResponse)

	// use custom method not allowed handler for the httprouter method not allowed
	router.MethodNotAllowed = http.HandlerFunc(app.notAllowedResponse)

	// register endpoints
	router.HandlerFunc(http.MethodGet, "/v1/healthcheck", app.healthcheckHandler)
	router.HandlerFunc(http.MethodPost, "/v1/movies", app.createMovieHandler)
	router.HandlerFunc(http.MethodGet, "/v1/movies/:id", app.showMovieHandler)
	return app.recoverPanic(router)
}
