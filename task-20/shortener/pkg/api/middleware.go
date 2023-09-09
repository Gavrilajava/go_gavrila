package api

import (
	"context"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/Gavrilajava/go_gavrila/task-20/common"
)

type contextKey string

var appMetrics = common.New()

func requestIDMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), contextKey("requestId"), uuid.New())
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func requestTimeOutMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithTimeout(r.Context(), time.Second*10)
		defer cancel()
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func setResponseHeaders(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

func prometheusMiddleware(next http.Handler) http.Handler {
  return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    route := mux.CurrentRoute(r)
    path, _ := route.GetPathTemplate()
    timer := prometheus.NewTimer(appMetrics.ResponseTime.WithLabelValues(path))
    next.ServeHTTP(w, r)
    timer.ObserveDuration()
		appMetrics.RequestsTotal.WithLabelValues(path).Inc()
  })
}

func logMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		common.Log(r, "started")
		next.ServeHTTP(w, r)
		common.Log(r, "finished")
	})
}
