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
var postgresUser, postgresPassword, postgresProject, postgresRegion, postgresInstance, postgresDB string

func init() {
	flag.UintVar(&port, "port", 8080, "port for the server to listen on")
	flag.StringVar(&postgresUser, "postgres_user", getenvWithDefault("GCP_POSTGRES_USER", "postgres"),
		"user for the postgres database")
	flag.StringVar(&postgresPassword, "postgres_password", getenvWithDefault("GCP_POSTGRES_PASSWORD", ""), "password for the postgres user")
	flag.StringVar(&postgresProject, "postgres_project", getenvWithDefault("GCP_POSTGRES_PROJECT", "tensile-imprint-156310"), "GCP project the GCP instance is in")
	flag.StringVar(&postgresRegion, "postgres_region", getenvWithDefault("GCP_POSTGRES_REGION", "europe-west1"), "region the postgres database is in")
	flag.StringVar(&postgresInstance, "postgres_instance", getenvWithDefault("GCP_POSTGRES_INSTANCE", "house-app"), "postgres database instance to query")
	flag.StringVar(&postgresDB, "postgres_db", getenvWithDefault("GCP_POSTGRES_DB_NAME", "postgres"), "name of the postgres database to query")
}

func main() {
	flag.Parse()

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

func getenvWithDefault(env, def string) string {
	val := os.Getenv(env)
	if val != "" {
		return val
	}
	return def
}

func makeURLFromEnvVars() (*url.URL, error) {
	if postgresUser == "" {
		return nil, fmt.Errorf("no user specified")
	}
	if postgresPassword == "" {
		return nil, fmt.Errorf("no user password specified")
	}
	if postgresProject == "" {
		return nil, fmt.Errorf("no project specified")
	}
	if postgresRegion == "" {
		return nil, fmt.Errorf("no region specified")
	}
	if postgresInstance == "" {
		return nil, fmt.Errorf("no intsance specified")
	}
	if postgresDB == "" {
		return nil, fmt.Errorf("no database name specified")
	}
	return &url.URL{
		Scheme: "gcppostgres",
		User:   url.UserPassword(postgresUser, postgresPassword),
		Host:   postgresProject,
		Path:   fmt.Sprintf("%s/%s/%s", postgresRegion, postgresInstance, postgresDB),
	}, nil
}
