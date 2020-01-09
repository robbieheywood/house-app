package main

import (
	"fmt"
	"net/http"

	"go.uber.org/zap"

	"github.com/go-chi/chi"
)

type HouseServer struct {
	*chi.Mux
	logger       *zap.SugaredLogger
	authEndpoint string
	port         uint
}

func New(authEndpoint string, port uint, logger *zap.SugaredLogger) HouseServer {
	srv := HouseServer{
		Mux:          chi.NewRouter(),
		logger:       logger,
		authEndpoint: authEndpoint,
		port:         port,
	}
	srv.Get("/", srv.handle)
	return srv
}

func (srv *HouseServer) ListenAndServe() error {
	return http.ListenAndServe(fmt.Sprintf(":%v", srv.port), srv)
}

// handle does the actual handling of HTTP requests
func (srv *HouseServer) handle(w http.ResponseWriter, r *http.Request) {
	// The logging here is deliberately overzealous to work the logging pipeline
	srv.logger.Info("Handling request")

	// Auth check is ridiculously basic so don't need to extract password
	user, _, ok := r.BasicAuth()
	if !ok {
		http.Error(w, "no user found in request", http.StatusBadRequest)
		return
	}
	srv.logger.Infof("Extracted user %v from request", user)

	// For now, all we do is check that the user exists with the auth endpoint
	resp, err := http.Get(srv.authEndpoint + user)
	if err != nil {
		http.Error(w, "error authorizing user", http.StatusUnauthorized)
		return
	} else if !wasSuccessful(resp) {
		http.Error(w, "user not authorized", resp.StatusCode)
		return
	}
	srv.logger.Infof("User %v is authorized", user)

	if _, err := w.Write([]byte("Hello, world")); err != nil {
		http.Error(w, "error writing response", http.StatusInternalServerError)
	}
	srv.logger.Info("Finished handling request")
}

func wasSuccessful(resp *http.Response) bool {
	return resp.StatusCode >= 200 && resp.StatusCode < 300
}
