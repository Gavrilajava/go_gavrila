package api

import (
	"encoding/json"
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/gorilla/mux"
	"github.com/Gavrilajava/go_gavrila/task-20/common"

	"shortener/pkg/cache"
	"shortener/pkg/storage"
)

type Api struct {
	db 			*storage.DB
	cache   *cache.Service
	router 	*mux.Router
}

func New(db *storage.DB, cache *cache.Service) *Api {
	api := Api{
		db:   db,
		cache: cache,
		router: mux.NewRouter(),
	}
	api.router.Use(requestIDMiddleware)
	api.router.Use(requestTimeOutMiddleware)
	api.router.Use(setResponseHeaders)

	api.router.HandleFunc("/", api.add).Methods("POST")
	api.router.HandleFunc("/{short}", api.get).Methods("GET")
	api.router.Handle("/metrics", promhttp.Handler())

	return &api
}

func (api *Api) Router() *mux.Router {
	return api.router
}


func (api *Api) add(w http.ResponseWriter, r *http.Request) {

	record := storage.Record{}

	err := json.NewDecoder(r.Body).Decode(&record)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		common.Error(r, err)
		return
	}

	record, err = api.db.Add(r.Context(), record.Long)
	if err != nil {
		http.Error(w, "Error creating short URL", http.StatusInternalServerError)
		common.Error(r, err)
		return
	}

	err = api.cache.Add(r.Context(), record)
	if err != nil {
		common.Error(r, err)
	}

	json.NewEncoder(w).Encode(record)

}

func (api *Api) get(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	short := vars["short"]


	record, err := api.cache.Get(r.Context(), short)
	if err != nil {
		http.Redirect(w, r, record.Long, http.StatusSeeOther)
		return
	}

	record, err = api.db.Get(r.Context(), short)
	if err != nil {
		common.Error(r, err)
		http.Error(w, "URL not found", http.StatusNotFound)
		return
	}

	http.Redirect(w, r, record.Long, http.StatusSeeOther)

}