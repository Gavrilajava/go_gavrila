package main

import (
	"log"
	"net/http"
	"os"

	"github.com/Gavrilajava/go_gavrila/task-20/common"

	"shortener/pkg/api"
	"shortener/pkg/cache"
	"shortener/pkg/storage"
)


func main() {

	db_name := os.Getenv("POSTGRES_DB")
	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")


	db, err := storage.New("localhost", 5342, db_name, user, password, "disable", 10) 
	if err != nil {
		log.Fatalf("failed to initialize database: %v", err)
	}

	c := cache.New("http://localhost:8081/")

	api := api.New(db, c)

	common.Log(nil, "Server is running on port 8080")
	http.ListenAndServe(":8080", api.Router())
}