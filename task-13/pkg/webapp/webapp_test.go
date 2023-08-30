package webapp

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gorilla/mux"

	"go-gavrila/task-12/pkg/crawler"
	"go-gavrila/task-12/pkg/index"
)

var testMux *mux.Router

func TestMain(m *testing.M) {
	testMux = mux.NewRouter()
	testMux.HandleFunc("/index", indexHandler).Methods(http.MethodGet)
	testMux.HandleFunc("/docs", docsHandler).Methods(http.MethodGet)
	data = index.Service{}
	data.Index = make(map[string][]int)
	data.Index["go"] = []int{1, 2, 3, 4}
	data.Documents = []crawler.Document{
		{ID: 1, URL: "https://go.dev/solutions/americanexpress", Title: "First Site"},
		{ID: 2, URL: "https://go.dev/learn#featured-books", Title: "Second Site"},
	}
	m.Run()
}

func Test_indexHandler(t *testing.T) {

	payload, _ := json.Marshal(data)

	req := httptest.NewRequest(http.MethodPost, "/index", bytes.NewBuffer(payload))
	req.Header.Add("Content-type", "plain/text")
	rr := httptest.NewRecorder()
	testMux.ServeHTTP(rr, req)

	if rr.Code != http.StatusMethodNotAllowed {
		t.Errorf("error: got %d, want %d", rr.Code, http.StatusOK)
	}

	req = httptest.NewRequest(http.MethodGet, "/index", nil)
	req.Header.Add("Content-type", "plain/text")
	rr = httptest.NewRecorder()
	testMux.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("error: got %d, want %d", rr.Code, http.StatusOK)
	}
	got := rr.Body.String()
	want, _ := json.Marshal(data.Index)
	if !strings.Contains(got, string(want)) {
		t.Errorf("error: got %s, want %s", got, want)
	}
}

func Test_docsHandler(t *testing.T) {

	payload, _ := json.Marshal(data)

	req := httptest.NewRequest(http.MethodPost, "/docs", bytes.NewBuffer(payload))
	req.Header.Add("Content-type", "plain/text")
	rr := httptest.NewRecorder()
	testMux.ServeHTTP(rr, req)

	if rr.Code != http.StatusMethodNotAllowed {
		t.Errorf("error: got %d, want %d", rr.Code, http.StatusOK)
	}

	req = httptest.NewRequest(http.MethodGet, "/docs", nil)
	req.Header.Add("Content-type", "plain/text")
	rr = httptest.NewRecorder()
	testMux.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("error: got %d, want %d", rr.Code, http.StatusOK)
	}
	got := rr.Body.String()
	want, _ := json.Marshal(data.Documents)
	if !strings.Contains(got, string(want)) {
		t.Errorf("error: got %s, want %s", got, want)
	}

}
