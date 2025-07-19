package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Application struct {
	Config Config
}

type Config struct {
	Address string
}

func (app *Application) mount() http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
	})

	return r
}

func (app *Application) Start() error {

	server := &http.Server{
		Addr:    app.Config.Address,
		Handler: app.mount(),
	}

	return server.ListenAndServe()
}
