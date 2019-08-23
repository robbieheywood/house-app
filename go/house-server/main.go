package main

import (
	"os"
)

func main() {
	srv := New()

	if err := srv.ListenAndServe(); err != nil {
		os.Exit(1)
	}
}
