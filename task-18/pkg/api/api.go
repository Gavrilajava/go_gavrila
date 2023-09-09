package api

import (
	"os"
	"fmt"
	"net/http"
	"sync"
	"encoding/json"
	"hash/fnv"

	"github.com/gorilla/mux"
)

type API struct {
	Router *mux.Router

	sync.Mutex
	db map[string]string
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
		db: make(map[string]string),
	}

	api.Router.HandleFunc("/{index}", api.getUrl).Methods(http.MethodGet)
	api.Router.HandleFunc("/", api.addUrl).Methods(http.MethodPost)

	return &api
}



func (a *API) getUrl(w http.ResponseWriter, r *http.Request) {
  index := mux.Vars(r)["index"]
	
	if long, ok := a.db[index]; ok {
		w.Write([]byte(fmt.Sprintf("%s", long)))
	}

	http.Error(w, "Resource Not Found", http.StatusNotFound)
}

func (a *API) addUrl(w http.ResponseWriter, r *http.Request) {
	params := &Params{}

	err := json.NewDecoder(r.Body).Decode(&params)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	hostname, _ := os.Hostname()

	short := hash(params.Url)
	a.db[short] = params.Url

	w.Write([]byte(fmt.Sprintf("%s/%s", hostname, short)))
}

func hash(s string) string {
	h := fnv.New32a()
	h.Write([]byte(s))
	return fmt.Sprintf("%d", h.Sum32())
}