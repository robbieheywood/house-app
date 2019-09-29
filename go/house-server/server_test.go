package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/require"
)

func TestChecksAuth_Correctly(t *testing.T) {
	tests := []struct {
		name string
		user string
	}{
		{name: "authoized user", user: "robbie"},
		{name: "unauthorized user", user: "wobbie"},
		{name: "empty user", user: ""},
	}

	fakeAuth := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r.URL.Path)
		if r.URL.Path != "/auth/robbie" {
			http.Error(w, "User not recognised", http.StatusUnauthorized)
		}
	}))
	defer fakeAuth.Close()
	fmt.Println(fakeAuth.URL)

	srv := HouseServer{
		rtr:          chi.NewRouter(),
		log:          logrus.New(),
		authEndpoint: fakeAuth.URL + "/auth/",
	}
	srv.rtr.Get("/", srv.handle)

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			req, err := http.NewRequest("GET", "/", nil)
			require.NoError(t, err)
			if test.user != "" {
				req.Header["User"] = []string{test.user}
			}

			w := httptest.NewRecorder()
			srv.handle(w, req)

			if test.user == "robbie" {
				require.Equal(t, w.Result().StatusCode, http.StatusOK)
			} else if test.user == "" {
				require.Equal(t, w.Result().StatusCode, http.StatusBadRequest)
			} else {
				require.Equal(t, w.Result().StatusCode, http.StatusUnauthorized)
			}
		})
	}
}
