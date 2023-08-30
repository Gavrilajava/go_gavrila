package api

import (
	"os"
	"testing"

	"go-gavrila/task-13/pkg/crawler"
	"go-gavrila/task-13/pkg/index"
)

var api *API

func TestMain(m *testing.M) {
	s := index.Service{
		Index:     make(map[string][]int),
		Documents: []crawler.Document{},
	}
	s.Add(crawler.Document{URL: "https://google.com", Title: "Google"})
	s.Add(crawler.Document{URL: "https://yandex.ru", Title: "Yandex"})
	api = New(&s)
	os.Exit(m.Run())
}
