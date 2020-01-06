package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/url"
	"os"

	"github.com/sirupsen/logrus"

	"gocloud.dev/postgres"
	_ "gocloud.dev/postgres/gcppostgres"
)

var port uint

func init() {
	flag.UintVar(&port, "port", 8080, "port for the server to listen on")
}

func main() {
	dbURL, err := makeURLFromEnvVars()
	if err != nil {
		log.Fatalf("failed to obtain URL for database: %v", err)
	}

	db, err := postgres.Open(context.Background(), dbURL.String())
	if err != nil {
		log.Fatalf("failed to open connection to database: %v", err)
	}

	// Ensure the database connection works
	err = db.Ping()
	if err != nil {
		log.Fatalf("failed to run initial query against database: %v", err)
	}

	srv := New(db, port, logrus.New())

	if err := srv.ListenAndServe(); err != nil {
		log.Fatalf("server error: %v", err)
	}
}

func makeURLFromEnvVars() (*url.URL, error) {
	user := os.Getenv("GCP_POSTGRES_USER")
	if user == "" {
		return nil, fmt.Errorf("no user specified")
	}
	password := os.Getenv("GCP_POSTGRES_PASSWORD")
	if user == "" {
		return nil, fmt.Errorf("no user password specified")
	}
	project := os.Getenv("GCP_POSTGRES_PROJECT")
	if user == "" {
		return nil, fmt.Errorf("no project specified")
	}
	region := os.Getenv("GCP_POSTGRES_REGION")
	if user == "" {
		return nil, fmt.Errorf("no region specified")
	}
	instanceName := os.Getenv("GCP_POSTGRES_INSTANCE")
	if user == "" {
		return nil, fmt.Errorf("no intsance specified")
	}
	dbName := os.Getenv("GCP_POSTGRES_DB_NAME")
	if user == "" {
		return nil, fmt.Errorf("no database name specified")
	}
	return &url.URL{
		Scheme: "gcppostgres",
		User:   url.UserPassword(user, password),
		Host:   project,
		Path:   fmt.Sprintf("%s/%s/%s", region, instanceName, dbName),
	}, nil
}
