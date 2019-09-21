package main

import (
	"flag"
	"os"
)

var authEndpoint string

func init() {
	flag.StringVar(&authEndpoint, "auth endpoint address", "https://tensile-imprint-156310.appspot.com/auth/",
		"address of the auth endpoint to hit")
}

func main() {
	srv := New(authEndpoint)

	if err := srv.ListenAndServe(); err != nil {
		os.Exit(1)
	}
}
