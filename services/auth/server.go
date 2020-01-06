package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/sirupsen/logrus"
)

type AuthServer struct {
	*chi.Mux
	log  *logrus.Logger
	db   *sql.DB
	port uint
}

// New creates a new house-server
func New(db *sql.DB, port uint, logger *logrus.Logger) AuthServer {
	app := AuthServer{
		Mux:  chi.NewRouter(),
		log:  logger,
		db:   db,
		port: port,
	}
	app.Get("/auth/*", app.authUser)

	return app
}

func (srv *AuthServer) ListenAndServe() error {
	return http.ListenAndServe(fmt.Sprintf(":%v", srv.port), srv)
}

// authUser does the actual handling of authorization requests
func (srv *AuthServer) authUser(w http.ResponseWriter, r *http.Request) {
	reqUser := r.URL.Path[len("/auth/"):]

	if err := srv.checkUser(context.Background(), reqUser); err != nil {
		http.Error(w, "User not recognised", http.StatusUnauthorized)
	}

	return
}

// checkUser looks up users in the database and checks whether they exist or not
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
