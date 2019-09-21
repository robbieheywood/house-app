package main

import (
	"fmt"
	"github.com/go-chi/chi"
	"github.com/sirupsen/logrus"
	"net/http"
)

// HouseServer defines the house-server used in this package
type HouseServer struct {
	rtr *chi.Mux
	log *logrus.Logger
	callAuthEndpoint func(user string) (resp *http.Response, err error)
}

// New creates a new house-server
func New(authEndpoint string) HouseServer {
	callAuthEndpoint := func(user string) (resp *http.Response, err error) {
		return http.Get(authEndpoint + user)
	}

	app := HouseServer{
		rtr: chi.NewRouter(),
		log: logrus.New(),
		callAuthEndpoint: callAuthEndpoint,
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
	user, ok := r.Header["User"]
	if !ok {
		http.Error(w, "no user found in request", http.StatusBadRequest)
		return
	} else if len(user) != 1 {
		http.Error(w, "request must only specify 1 user", http.StatusBadRequest)
		return
	}

	fmt.Println(user)
	resp, err := srv.callAuthEndpoint(user[0])
	fmt.Println(resp.StatusCode)
	fmt.Println(err)
	if err != nil {
		http.Error(w, "error authorizing user", http.StatusUnauthorized)
		return
	} else if !hasSuccessStatus(resp) {
		http.Error(w, "user not authorized", resp.StatusCode)
		return
	}

	if _, err := w.Write([]byte("Hello, world")); err != nil {
		http.Error(w, "error writing response", http.StatusInternalServerError)
		return
		srv.log.Errorf("error writing logs: %v", err)
	}
}

// hasSuccessStatus checks if the http Response has a 'success' status or not
func hasSuccessStatus(resp *http.Response) bool {
	return resp.StatusCode >= 200 && resp.StatusCode < 300
}