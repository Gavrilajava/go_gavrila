package common

import (
	"os"
	"fmt"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/rs/zerolog"
)

type Metrics struct {
	ResponseTime      *prometheus.HistogramVec
	RequestsTotal *prometheus.CounterVec
}

var log = zerolog.New(zerolog.ConsoleWriter{Out: os.Stdout}).With().Timestamp().Logger()

type contextKey string

func New() *Metrics {
	reg := prometheus.NewRegistry()
	m := &Metrics{
		ResponseTime: prometheus.NewHistogramVec(prometheus.HistogramOpts{
			Name: "response_time",
			Help: "Server response time",
		}, []string{"route"}),
		RequestsTotal: prometheus.NewCounterVec(prometheus.CounterOpts{
			Name: "requests_total",
			Help: "Number of requests.",
		}, []string{"route"}),
	}

	reg.MustRegister(m.ResponseTime)
	reg.MustRegister(m.RequestsTotal)

	return m
}



func Log(r *http.Request, message string) {
	if r != nil {
		log.Info().Str("requestInfo", requestInfo(r)).Msg(message)
	} else {
		log.Info().Msg(message)
	}
}

func Error(r *http.Request, err error) {
	if r != nil {
		log.Error().Str("requestInfo", requestInfo(r)).Msg(err.Error())
	} else {
		log.Info().Msg(err.Error())
	}
}


func requestInfo(r *http.Request) string {
	return fmt.Sprintf("%s %s %s %s", r.Method, r.RemoteAddr, r.RequestURI, r.Context().Value(contextKey("requestId")))
}