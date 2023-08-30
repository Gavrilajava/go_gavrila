package api

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
)

type contextKey string

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

func logMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.Method, r.RemoteAddr, r.RequestURI, r.Context().Value(contextKey("requestId")))
		next.ServeHTTP(w, r)
	})
}
