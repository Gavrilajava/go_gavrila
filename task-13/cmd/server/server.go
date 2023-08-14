package main

import (
	"log"
	"net/http"

	"go-gavrila/task-13/pkg/api"
	"go-gavrila/task-13/pkg/index"
)

func main() {

	index, err := index.New()
	if err != nil {
		log.Fatal(err)
	}

	api := api.New(index)

	http.ListenAndServe(":8082", api.Router())

}
