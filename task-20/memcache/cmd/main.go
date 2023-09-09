package main

import (
	"log"
	"net/http"
	"os"
	"metrics/pkg/metrics"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/rs/zerolog"
	"github.com/gorilla/mux"

	"memcache/pkg/api"
	"memcache/pkg/redis"
)

func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	l := zerolog.New(zerolog.ConsoleWriter{Out: os.Stdout}).With().Timestamp().Logger()

	reg := prometheus.NewRegistry()
	m := metrics.New(reg)

	redisStorage, err := redis.New("localhost:6379", "", 0)
	if err != nil {
		log.Fatalf("Failed to initialize Redis storage: %v", err)
	}

	api := api.New(redisStorage, m, l)

	r := mux.NewRouter()
	api.Init(r)

	r.Handle("/metrics", promhttp.Handler())

	log.Println("Server is running on http://localhost:8081")
	log.Fatal(http.ListenAndServe(":8081", r))
}