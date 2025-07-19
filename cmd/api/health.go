package main

import "net/http"

func (app *Application) healthCheck(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Everything is OK"))
}
