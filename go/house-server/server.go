package main

import (
	"github.com/go-chi/chi"
	"github.com/sirupsen/logrus"
	"net/http"
)

type App struct {
	rtr *chi.Mux
	log *logrus.Logger
}

func New() App {
	rtr := chi.NewRouter()
	app := App{
		rtr: rtr,
		log: logrus.New(),
	}
	app.rtr.Get("/", app.handle)

	return app
}

func (app *App) ListenAndServe() error {
	return http.ListenAndServe(":8080", app.rtr)
}

func (app* App) handle(w http.ResponseWriter, r *http.Request) {
	if _, err := w.Write([]byte("Hello, world")); err != nil {
		app.log.Errorf("error writing logs: %v", err)
	}
}
