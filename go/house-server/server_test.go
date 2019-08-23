package main

import (
	"github.com/stretchr/testify/require"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestApp_HandlesBaselineRequests_Successfully(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	require.NoError(t, err)
	w := httptest.NewRecorder()

	app := New()
	app.handle(w, req)

	require.Equal(t, w.Result().StatusCode, http.StatusOK)
	body, err := ioutil.ReadAll(w.Result().Body)
	require.NoError(t, err)
	require.Equal(t, body, []byte("Hello, world"))
}
