package main

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestUsersAreAuthed_Correctly(t *testing.T) {
	tests := []struct {
		name       string
		user       string
		expectPass bool
	}{
		{name: "authorized user", user: "robbie", expectPass: true},
		{name: "unauthorized user - capitalization", user: "Robbie", expectPass: false},
		{name: "unathorized user - different name", user: "wobbie", expectPass: false},
		{name: "empty user", user: "", expectPass: false},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			path := fmt.Sprintf("/auth/%v", test.user)
			req, err := http.NewRequest("GET", path, nil)
			require.NoError(t, err)
			w := httptest.NewRecorder()

			srv := New()
			srv.authUser(w, req)

			if test.expectPass {
				require.Equal(t, w.Result().StatusCode, http.StatusOK)
			} else {
				require.Equal(t, w.Result().StatusCode, http.StatusUnauthorized)
			}
		})
	}
}
