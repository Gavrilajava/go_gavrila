package main

import (
	"net/http"

	"shortener/pkg/api"
)

func main() {
	srv := New()
	srv.run()
}

type Server struct {
	api *api.API
}

func New() *Server {
	s := Server{
		api: api.New(),
	}
	return &s
}

func (s *Server) run() {
	http.ListenAndServe(":8080", s.api.Router)
}