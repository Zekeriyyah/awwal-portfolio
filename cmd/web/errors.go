package main

import (
	"fmt"
	"net/http"
)

func (app *App) clientErr(w http.ResponseWriter, err error, c int) {
	app.errorLog.Println(err)
	fmt.Fprint(w, c, http.StatusText(c))
}
