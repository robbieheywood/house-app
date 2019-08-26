package main

import (
	"github.com/go-chi/chi"
	"github.com/sirupsen/logrus"
	"net/http"
)

// HouseServer defines the house-server used in this package
type HouseServer struct {
	rtr *chi.Mux
	log *logrus.Logger
}

// New creates a new house-server
func New() HouseServer {
	app := HouseServer{
		rtr: chi.NewRouter(),
		log: logrus.New(),
	}
	app.rtr.Get("/", app.handle)

	return app
}

// ListenAndServe starts the server. It always returns a non-nil error
func (srv *HouseServer) ListenAndServe() error {
	return http.ListenAndServe(":8080", srv.rtr)
}

// handle does the actual handling of HTTP requests
func (srv *HouseServer) handle(w http.ResponseWriter, r *http.Request) {
	if _, err := w.Write([]byte("Hello, world")); err != nil {
		srv.log.Errorf("error writing logs: %v", err)
	}
}
