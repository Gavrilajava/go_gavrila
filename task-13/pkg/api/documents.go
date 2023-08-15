package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"go-gavrila/task-13/pkg/crawler"
)

func (api *API) documents(w http.ResponseWriter, r *http.Request) {
	err := json.NewEncoder(w).Encode(api.index.Documents)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (api *API) document(w http.ResponseWriter, r *http.Request) {
	d, code, err := api.getDocument(r)
	if err != nil {
		http.Error(w, err.Error(), code)
		return
	}
	err = json.NewEncoder(w).Encode(d)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (api *API) createDocument(w http.ResponseWriter, r *http.Request) {
	params, err := documentParams(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}
	d := api.index.Add(params)
	api.index.Save()
	err = json.NewEncoder(w).Encode(d)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}

func (api *API) updateDocument(w http.ResponseWriter, r *http.Request) {
	d, code, err := api.getDocument(r)
	if err != nil {
		http.Error(w, err.Error(), code)
		return
	}
	params, err := documentParams(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	d = api.index.Update(d, params)
	api.index.Save()
	err = json.NewEncoder(w).Encode(d)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (api *API) destroyDocument(w http.ResponseWriter, r *http.Request) {
	d, code, err := api.getDocument(r)
	if err != nil {
		http.Error(w, err.Error(), code)
		return
	}
	api.index.Delete(d)
	err = json.NewEncoder(w).Encode(api.index.Documents)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (api *API) getDocument(r *http.Request) (crawler.Document, int, error) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		return crawler.Document{}, http.StatusUnprocessableEntity, err
	}
	d, err := api.index.Find(id)
	if err != nil {
		return crawler.Document{}, http.StatusNotFound, err
	}
	return d, 0, nil
}

func documentParams(r *http.Request) (crawler.Document, error) {
	var d crawler.Document
	err := json.NewDecoder(r.Body).Decode(&d)
	if err != nil {
		return crawler.Document{}, err
	}
	return d, nil
}
