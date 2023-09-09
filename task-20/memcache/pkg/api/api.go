package api

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/Gavrilajava/go_gavrila/task-20/common"

	"memcache/pkg/redis"
)


type Api struct {
	redis   *redis.Storage
	router *mux.Router
}


func New(redis *redis.Storage) *Api {
	api := Api{
		redis:   redis,
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

	var record redis.Record

	err := json.NewDecoder(r.Body).Decode(&record)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		common.Error(r, err)
		return
	}

	err = api.redis.Add(record)
	if err != nil {
		http.Error(w, "Error creating short URL", http.StatusInternalServerError)
		common.Error(r, err)
		return
	}

	json.NewEncoder(w).Encode(record)

}

func (api *Api) get(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	short := vars["short"]

	record, err := api.redis.Get(short)
	if err != nil {
		common.Error(r, err)
		http.Error(w, "URL not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(record)

}