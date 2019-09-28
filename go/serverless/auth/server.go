package main

import (
	"context"
	"database/sql"
	"errors"
	"github.com/go-chi/chi"
	"github.com/sirupsen/logrus"
	"net/http"
	"time"
)

// HouseServer defines the house-server used in this package
type AuthServer struct {
	rtr *chi.Mux
	log *logrus.Logger
	db  *sql.DB
}

// New creates a new house-server
func New(db *sql.DB) AuthServer {
	app := AuthServer{
		rtr: chi.NewRouter(),
		log: logrus.New(),
		db: db,
	}
	app.rtr.Get("/auth/*", app.authUser)

	return app
}

// ListenAndServe starts the server. It always returns a non-nil error
func (srv *AuthServer) ListenAndServe() error {
	return http.ListenAndServe(":8080", srv.rtr)
}

// handle does the actual handling of HTTP requests
func (srv *AuthServer) authUser(w http.ResponseWriter, r *http.Request) {
	reqUser := r.URL.Path[len("/auth/"):]

	if err := srv.checkUser(context.Background(), reqUser); err != nil {
		srv.log.Errorf("failed to check user: %v", err)
		http.Error(w, "User not recognised", http.StatusUnauthorized)
	}

	return
}

// checkUser looks up users in the database and checks whether they are valid
func (srv *AuthServer) checkUser(ctx context.Context, user string) error {
	cmd := "SELECT name FROM users WHERE name=$1"
	ctx, cancel := context.WithTimeout(ctx, time.Minute)
	defer cancel()
	rows, err := srv.db.QueryContext(ctx, cmd, user)
	if err != nil {
		return errors.New("failed to check user with database: " + err.Error())
	}
	defer rows.Close()

	if !rows.Next() {
		return errors.New("user not found")
	}

	return nil
}
