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
		name string
		expectPass bool
	}{
		{"robbie", true},
		{"Robbie", false},
		{"wobbie", false},
		{"", false},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			path := fmt.Sprintf("/auth/%v", test.name)
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

