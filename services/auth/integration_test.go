package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/cenkalti/backoff"
	"github.com/sirupsen/logrus"

	_ "github.com/lib/pq"
	"github.com/ory/dockertest"
	"github.com/stretchr/testify/require"
)

const dbName = "test-db"
const userName = "test-user"
const password = "test-password"

func TestUsersAreAuthed_Correctly(t *testing.T) {
	tests := []struct {
		name       string
		expectPass bool
	}{
		{"robbie", true},
		{"Robbie", false},
		{"wobbie", false},
		{"", false},
	}
	port, cleanup := createTestDB(t)
	defer cleanup()
	db := setupTestDB(t, port)

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			req, err := http.NewRequest("GET", fmt.Sprintf("/auth/%v", test.name), nil)
			require.NoError(t, err)
			w := httptest.NewRecorder()
			srv := New(db, 8080, logrus.New())
			srv.authUser(w, req)

			if test.expectPass {
				require.Equal(t, w.Result().StatusCode, http.StatusOK)
			} else {
				require.Equal(t, w.Result().StatusCode, http.StatusUnauthorized)
			}
		})
	}
}

// createTestDB starts running a postgres database instance in a docker container for the test
func createTestDB(t *testing.T) (string, func()) {
	pool, err := dockertest.NewPool("")
	require.NoError(t, err)
	envVars := []string{
		"POSTGRES_USER=" + userName,
		"POSTGRES_PASSWORD=" + password,
		"POSTGRES_DB=" + dbName,
	}
	resource, err := pool.Run("postgres", "11", envVars)
	require.NoError(t, err)
	cleanup := func() {
		resource.Close()
	}
	return resource.GetPort("5432/tcp"), cleanup
}

// setuptestDB creates the connection to the test database and sets it up with test data
func setupTestDB(t *testing.T, port string) *sql.DB {
	// Wait for the DB to be ready
	var db *sql.DB
	err := backoff.Retry(func() error {
		var err error
		db, err = sql.Open("postgres",
			fmt.Sprintf("user=%v password=%v dbname=%v port=%v sslmode=disable", userName, password, dbName, port))
		if err != nil {
			return err
		}
		return db.Ping()
	}, backoff.WithMaxRetries(backoff.NewExponentialBackOff(), 5))
	require.NoError(t, err)

	_, err = db.Exec("CREATE Table users (name VARCHAR(255), id SERIAL PRIMARY KEY)")
	require.NoError(t, err)
	_, err = db.Exec("INSERT INTO users (name) values ('robbie')")
	require.NoError(t, err)

	return db
}
