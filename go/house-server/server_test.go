package main

import (
	"github.com/go-chi/chi"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"testing"
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

	req, err := http.NewRequest("GET", "/", nil)
	require.NoError(t, err)

	fakeAuthCall := func(user string) (*http.Response, error) {
		if user == "robbie" {
			return &http.Response{StatusCode: http.StatusOK}, nil
		} else {
			return &http.Response{StatusCode: http.StatusUnauthorized}, nil
		}
	}

	srv := HouseServer{
		rtr:              chi.NewRouter(),
		log:              logrus.New(),
		callAuthEndpoint: fakeAuthCall,
	}
	srv.rtr.Get("/", srv.handle)

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			if test.user != "" {
				req.Header["User"] = []string{test.user}
			}

			srv.handle(w, req)

			if test.user == "robbie" {
				require.Equal(t, w.Result().StatusCode, http.StatusOK)
			} else {
				require.Equal(t, w.Result().StatusCode, http.StatusUnauthorized)
			}
		})
	}
}
