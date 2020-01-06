package main

import (
	"flag"
	"log"

	"go.uber.org/zap"
)

var authEndpoint string
var port uint

func init() {
	flag.StringVar(&authEndpoint, "auth endpoint address", "https://tensile-imprint-156310.appspot.com/auth/",
		"address of the auth endpoint to hit")
	flag.UintVar(&port, "port", 8080, "port for the server to listen on")
}

func main() {
	logger, err := zap.NewDevelopment()
	if err != nil {
		log.Fatalf("can't initialize zap logger: %v", err)
	}
	sugaredLogger := logger.Sugar()
	defer func() {
		_ = logger.Sync()
	}()

	srv := New(authEndpoint, port, sugaredLogger)

	if err := srv.ListenAndServe(); err != nil {
		log.Fatalf("server error: %v", err)
	}
}
