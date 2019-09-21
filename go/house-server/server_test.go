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
	}{
		{"robbie"},
		{"wobbie"},
		{""},
	}

	req, err := http.NewRequest("GET", "/", nil)
	require.NoError(t, err)
	w := httptest.NewRecorder()

	fakeAuthCall := func(user string) (*http.Response, error) {
		if user == "robbie" {
			return &http.Response{StatusCode: http.StatusOK}, nil
		} else {
			return &http.Response{StatusCode: http.StatusUnauthorized}, nil
		}
	}

	srv := HouseServer{
		rtr: chi.NewRouter(),
		log: logrus.New(),
		callAuthEndpoint: fakeAuthCall,
	}
	srv.rtr.Get("/", srv.handle)

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			srv.handle(w, req)

			if test.name == "robbie" {
				require.Equal(t, w.Result().StatusCode, http.StatusOK)
			} else {
				require.Equal(t, w.Result().StatusCode, http.StatusUnauthorized)
			}
		})
	}
}
