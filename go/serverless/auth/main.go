package main

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/GoogleCloudPlatform/cloudsql-proxy/proxy/dialers/postgres"
)

func main() {
	datastoreName := "user=postgres password=157AbbeyRoad dbname=postgres host=tensile-imprint-156310" +
		":europe-west1:house-users sslmode=disable" //os.Getenv("POSTGRES_CONNECTION")

	db, err := sql.Open("cloudsqlpostgres", datastoreName)
	if err != nil {
		log.Fatalf("failed to open connection to database: %v", err)
	}

	// Ensure the table exists.
	// Running an SQL query also checks the connection to the PostgreSQL server
	// is authenticated and valid.
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS users (name VARCHAR(255), entryid SERIAL PRIMARY KEY)")
	if err != nil {
		log.Fatalf("failed to run initial query against database: %v", err)
	}

	srv := New(db)

	if err := srv.ListenAndServe(); err != nil {
		os.Exit(1)
	}
}
