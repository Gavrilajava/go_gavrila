package main

import (
	"log"
	"net/http"

	"github.com/Gavrilajava/go_gavrila/task-20/common"

	"memcache/pkg/api"
	"memcache/pkg/redis"
)

func main() {

	redis, err := redis.New("localhost:6379", "", 0)
	if err != nil {
		log.Fatalf("Failed to initialize Redis storage: %v", err)
	}

	api := api.New(redis)

	common.Log(nil, "Server is running on port 8081")
	http.ListenAndServe(":8081", api.Router())
}