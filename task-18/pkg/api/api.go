package api

import (
	"os"
	"fmt"
	"net/http"
	"sync"
	"strconv"
	"encoding/json"

	"github.com/gorilla/mux"
)

type API struct {
	Router *mux.Router

	sync.Mutex
	db []string
}

type Link struct {
	short   string
	full    string
}

type Params struct {
	Url      string `json:"url"`
}

func New() *API {
	api := API{
		Router: mux.NewRouter(),
	}

	api.db = append(api.db, "www.google.com")

	api.Router.HandleFunc("/{index}", api.getUrl).Methods(http.MethodGet)
	api.Router.HandleFunc("/", api.addUrl).Methods(http.MethodPost)

	return &api
}



func (a *API) getUrl(w http.ResponseWriter, r *http.Request) {
  v := mux.Vars(r)["index"]

	index, err := strconv.Atoi(v)

	if err != nil || index < 0 ||  index >= len(a.db) {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Write([]byte(fmt.Sprintf("%s", a.db[index])))
}

func (a *API) addUrl(w http.ResponseWriter, r *http.Request) {
	params := &Params{}

	err := json.NewDecoder(r.Body).Decode(&params)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	hostname, _ := os.Hostname()

	w.Write([]byte(fmt.Sprintf("%s/%d", hostname, len(a.db) - 1)))
}