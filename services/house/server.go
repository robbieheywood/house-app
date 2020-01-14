package main

import (
	"fmt"
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/prometheus/client_golang/prometheus"

	"github.com/opentracing/opentracing-go"

	"go.uber.org/zap"

	"github.com/go-chi/chi"
)

var reqsHandled = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Name:      "reqs_handled_total",
		Subsystem: "house",
	},
	[]string{"code"})

func init() {
	prometheus.MustRegister(reqsHandled)
}

type HouseServer struct {
	*chi.Mux
	logger       *zap.SugaredLogger
	authEndpoint string
	port         uint
}

func New(authEndpoint string, port uint, logger *zap.SugaredLogger) HouseServer {
	srv := HouseServer{
		Mux:          chi.NewRouter(),
		logger:       logger,
		authEndpoint: authEndpoint,
		port:         port,
	}
	srv.Get("/metrics", promhttp.Handler().ServeHTTP)
	srv.Get("/", srv.handle)
	return srv
}

func (srv *HouseServer) ListenAndServe() error {
	return http.ListenAndServe(fmt.Sprintf(":%v", srv.port), srv)
}

// handle does the actual handling of HTTP requests
func (srv *HouseServer) handle(w http.ResponseWriter, r *http.Request) {
	// The logging here is deliberately overzealous to work the logging pipeline
	srv.logger.Info("Handling request")
	span := opentracing.StartSpan("house.handle")
	span.SetTag("http.method", r.Method)
	span.SetTag("http.url", r.URL.String())
	defer span.Finish()

	// Auth check is ridiculously basic so don't need to extract password
	user, _, ok := r.BasicAuth()
	if !ok {
		http.Error(w, "no user found in request", http.StatusBadRequest)
		span.SetTag("http.status", http.StatusBadRequest)
		reqsHandled.WithLabelValues(fmt.Sprintf("%v", http.StatusBadRequest)).Inc()
		return
	}
	srv.logger.Infof("Extracted user %v from request", user)

	// For now, all we do is check that the user exists with the auth endpoint
	authSpan := opentracing.StartSpan("auth.check", opentracing.ChildOf(span.Context()))
	authSpan.SetTag("auth.user", user)
	resp, err := http.Get(srv.authEndpoint + user)
	authSpan.Finish()
	if err != nil {
		http.Error(w, "error authorizing user", http.StatusUnauthorized)
		span.SetTag("http.status", http.StatusUnauthorized)
		reqsHandled.WithLabelValues(fmt.Sprintf("%v", http.StatusUnauthorized)).Inc()
		return
	} else if !wasSuccessful(resp) {
		http.Error(w, "user not authorized", resp.StatusCode)
		span.SetTag("http.status", resp.StatusCode)
		reqsHandled.WithLabelValues(fmt.Sprintf("%v", resp.StatusCode)).Inc()
		return
	}
	srv.logger.Infof("User %v is authorized", user)

	if _, err := w.Write([]byte("Hello, world")); err != nil {
		http.Error(w, "error writing response", http.StatusInternalServerError)
		span.SetTag("http.status", http.StatusInternalServerError)
		reqsHandled.WithLabelValues(fmt.Sprintf("%v", http.StatusInternalServerError)).Inc()
	}
	srv.logger.Info("Finished handling request")
	span.SetTag("http.status", http.StatusOK)
	reqsHandled.WithLabelValues(fmt.Sprintf("%v", http.StatusOK)).Inc()
}

func wasSuccessful(resp *http.Response) bool {
	return resp.StatusCode >= 200 && resp.StatusCode < 300
}
