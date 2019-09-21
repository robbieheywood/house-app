package main

import (
	"github.com/go-chi/chi"
	"github.com/sirupsen/logrus"
	"net/http"
)

// HouseServer defines the house-server used in this package
type AuthServer struct {
	rtr *chi.Mux
	log *logrus.Logger
}

var users = make(map[string]struct{})

func init() {
	users["robbie"] = struct{}{}
	users["adam"] = struct{}{}
	users["tom"] = struct{}{}
	users["sam"] = struct{}{}
	users["marina"] = struct{}{}
}

// New creates a new house-server
func New() AuthServer {
	app := AuthServer{
		rtr: chi.NewRouter(),
		log: logrus.New(),
	}
	app.rtr.Get("/auth/", app.authUser)

	return app
}

// ListenAndServe starts the server. It always returns a non-nil error
func (srv *AuthServer) ListenAndServe() error {
	return http.ListenAndServe(":8080", srv.rtr)
}

// handle does the actual handling of HTTP requests
func (srv *AuthServer) authUser(w http.ResponseWriter, r *http.Request) {
	reqUser := r.URL.Path[len("/auth/"):]

	if _, ok := users[reqUser]; !ok {
		http.Error(w, "User not recognised", http.StatusUnauthorized)
	}

	return
}