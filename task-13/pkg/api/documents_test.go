package api

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"go-gavrila/task-13/pkg/crawler"
)

func TestAPI_documents(t *testing.T) {

	req := httptest.NewRequest(http.MethodGet, "/api/v1/documents", nil)
	rr := httptest.NewRecorder()
	api.router.ServeHTTP(rr, req)
	t.Log("Response: ", rr.Body)
	if !(rr.Code == http.StatusOK) {
		t.Errorf("got: %d, want: %d", rr.Code, http.StatusOK)
	}

	want, _ := json.Marshal(api.index.Documents)
	got := strings.TrimSuffix(rr.Body.String(), "\n")

	if got != string(want) {
		t.Errorf("got: %s, want: %s", got, want)
	}

}

func TestAPI_document(t *testing.T) {

	req := httptest.NewRequest(http.MethodGet, "/api/v1/documents/2", nil)
	rr := httptest.NewRecorder()

	api.router.ServeHTTP(rr, req)

	if !(rr.Code == http.StatusOK) {
		t.Errorf("got: %d, want: %d", rr.Code, http.StatusOK)
	}

	want, _ := json.Marshal(api.index.Documents[1])
	got := strings.TrimSuffix(rr.Body.String(), "\n")

	if got != string(want) {
		t.Errorf("got: %s, want: %s", got, want)
	}

}

func TestAPI_createDocument(t *testing.T) {

	size := len(api.index.Documents)

	payload, _ := json.Marshal(crawler.Document{URL: "https://bing.com", Title: "Bing"})
	req := httptest.NewRequest(http.MethodPost, "/api/v1/documents", bytes.NewBuffer(payload))
	rr := httptest.NewRecorder()

	api.router.ServeHTTP(rr, req)

	if !(len(api.index.Documents) == size+1) {
		t.Errorf("got: %d, want: %d", len(api.index.Documents), size)
	}

	if !(rr.Code == http.StatusOK) {
		t.Errorf("got: %d, want: %d", rr.Code, http.StatusOK)
	}

	want, _ := json.Marshal(api.index.Documents[size])
	got := strings.TrimSuffix(rr.Body.String(), "\n")

	if got != string(want) {
		t.Errorf("got: %s, want: %s", got, want)
	}

}

func TestAPI_updateDocument(t *testing.T) {

	size := len(api.index.Documents)

	payload, _ := json.Marshal(crawler.Document{URL: "https://bing.com", Title: "Bing"})
	req := httptest.NewRequest(http.MethodPatch, "/api/v1/documents/1", bytes.NewBuffer(payload))
	rr := httptest.NewRecorder()

	api.router.ServeHTTP(rr, req)

	if !(len(api.index.Documents) == size) {
		t.Errorf("got: %d, want: %d", len(api.index.Documents), size)
	}

	if !(rr.Code == http.StatusOK) {
		t.Errorf("got: %d, want: %d", rr.Code, http.StatusOK)
	}

	want, _ := json.Marshal(crawler.Document{ID: 1, URL: "https://bing.com", Title: "Bing"})
	got := strings.TrimSuffix(rr.Body.String(), "\n")

	if got != string(want) {
		t.Errorf("got: %s, want: %s", got, want)

	}

}

func TestAPI_deleteDocument(t *testing.T) {

	size := len(api.index.Documents)

	req := httptest.NewRequest(http.MethodDelete, "/api/v1/documents/1", nil)
	rr := httptest.NewRecorder()

	api.router.ServeHTTP(rr, req)

	if !(len(api.index.Documents) == size-1) {
		t.Errorf("got: %d, want: %d", len(api.index.Documents), size)
	}

	if !(rr.Code == http.StatusOK) {
		t.Errorf("got: %d, want: %d", rr.Code, http.StatusOK)
	}

	want, _ := json.Marshal(api.index.Documents)
	got := strings.TrimSuffix(rr.Body.String(), "\n")

	if got != string(want) {
		t.Errorf("got: %s, want: %s", got, want)

	}

}
