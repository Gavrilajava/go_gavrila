package main

import (
	"log"

	"go-gavrila/task-12/pkg/crawler/spider"
	"go-gavrila/task-12/pkg/index"
	"go-gavrila/task-12/pkg/webapp"
)

const depth = 2

var urls = [2]string{"https://go.dev", "https://golang.org"}

func main() {

	index, err := index.New()
	if err != nil {
		log.Fatal(err)
	}

	if index.Empty() {
		s := spider.New()
		for _, url := range urls {
			res, err := s.Scan(url, depth)
			if err != nil {
				log.Fatal(err)
			}
			index.Add(res)
		}
		if err = index.Save(); err != nil {
			log.Fatal(err)
		}
	}

	err = webapp.Start(":8000", *index)
	if err != nil {
		log.Fatal(err)
	}
}
