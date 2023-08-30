package api

import (
	"net/http"

	"github.com/gorilla/mux"

	"go-gavrila/task-13/pkg/index"
)

type API struct {
	router *mux.Router
	index  *index.Service
}

func New(i *index.Service) *API {
	api := API{
		router: mux.NewRouter(),
		index:  i,
	}

	api.endpoints()

	return &api
}

func (api *API) Router() *mux.Router {
	return api.router
}

func (api *API) endpoints() {
	api.router.Use(requestIDMiddleware)
	api.router.Use(logMiddleware)
	api.router.Use(requestTimeOutMiddleware)

	api.router.HandleFunc("/api/v1/documents", api.documents).Methods(http.MethodGet, http.MethodOptions)
	api.router.HandleFunc("/api/v1/documents/{id}", api.document).Methods(http.MethodGet, http.MethodOptions)
	api.router.HandleFunc("/api/v1/documents", api.createDocument).Methods(http.MethodPost)
	api.router.HandleFunc("/api/v1/documents/{id}", api.updateDocument).Methods(http.MethodPatch)
	api.router.HandleFunc("/api/v1/documents/{id}", api.destroyDocument).Methods(http.MethodDelete)

}
