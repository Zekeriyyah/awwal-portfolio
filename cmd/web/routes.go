package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *App) Route() http.Handler {
	router := httprouter.New()

	router.GET("/api/home/:mode", app.homeHandler)
	// 	// router.GET("/api/about/:mode", app.aboutHandler)
	// 	// router.GET("/api/projects/:mode", app.getProjects)
	// 	// router.GET("/api/projects/:mode", app.getProjectsByFilter)

	return router
}
