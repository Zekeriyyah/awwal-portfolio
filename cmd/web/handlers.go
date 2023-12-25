package main

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *App) homeHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	// 	params := httprouter.ParamsFromContext(r.Context())

	// 	id, err := strconv.Atoi(params.ByName("mode"))
	// 	if err != nil {
	// 		app.clientErr(w, err, http.StatusBadRequest)
	// 		return
	// 	}
	fmt.Fprint(w, "Fetch on the data needed to implement home page of the portfolio")
}
