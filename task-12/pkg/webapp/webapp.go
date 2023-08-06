package webapp

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"go-gavrila/task-12/pkg/index"
)

var data index.Service

func Start(port string, i index.Service) {
	data = i
	mux := mux.NewRouter()
	mux.HandleFunc("/index", indexHandler).Methods(http.MethodGet)
	mux.HandleFunc("/docs", docsHandler).Methods(http.MethodGet)
	fmt.Println("Starting server on port: " + port)
	err := http.ListenAndServe(port, mux)
	if err != nil {
		log.Fatal(err)
	}

}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	err := json.NewEncoder(w).Encode(data.Index)
	if err != nil {
		log.Print(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func docsHandler(w http.ResponseWriter, r *http.Request) {
	err := json.NewEncoder(w).Encode(data.Documents)
	if err != nil {
		log.Print(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
