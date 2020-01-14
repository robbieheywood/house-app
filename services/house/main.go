package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	jaegercfg "github.com/uber/jaeger-client-go/config"
	jaegerlog "github.com/uber/jaeger-client-go/log"
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

	fmt.Println("JAEGER_AGENT_HOST: " + os.Getenv("JAEGER_AGENT_HOST"))
	fmt.Println("JAEGER_AGENT_PORT: " + os.Getenv("JAEGER_AGENT_PORT"))

	cfg := &jaegercfg.Configuration{
		ServiceName: "house",
		Sampler: &jaegercfg.SamplerConfig{
			Type:  jaeger.SamplerTypeConst,
			Param: 1,
		},
		Reporter: &jaegercfg.ReporterConfig{
			LogSpans: true,
		},
	}
	cfg, err = cfg.FromEnv()
	if err != nil {
		log.Fatalf("error initialising Jaeger config: %v", err)
	}

	tracer, closer, err := cfg.NewTracer(jaegercfg.Logger(jaegerlog.StdLogger))
	if err != nil {
		log.Fatalf("error initialising Jaeger tracer: %v", err)
	}
	opentracing.SetGlobalTracer(tracer)
	defer closer.Close()

	srv := New(authEndpoint, port, sugaredLogger)

	if err := srv.ListenAndServe(); err != nil {
		log.Fatalf("server error: %v", err)
	}
}
