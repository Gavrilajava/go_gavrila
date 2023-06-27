package main

import (
	"go-gavrila/task-2/pkg/crawler"
	"go-gavrila/task-2/pkg/crawler/spider"
	"strings"
	"flag"
	"fmt"
)

const depth = 3
var urls = [2]string{"https://go.dev", "https://golang.org"}

func main() {

	s := spider.New()
	var docs []crawler.Document
	token := flag.String("s", "", "search string")

	for _, url := range urls {
		res, err := s.Scan(url, depth)
		if err != nil {
			fmt.Println(err)
			return
		} else {
			docs = append(docs, res...)
		}
	}

	flag.Parse()

	if *token != "" {
		fmt.Println("Search Results:")
		var f bool
		for _, d := range docs {
			if strings.Contains(strings.ToLower(d.Title), strings.ToLower(*token)) {
				if !f {
					f = true
				}
				fmt.Println(d.URL, d.Title)
			}
		}
		if !f {
			fmt.Println("Nothing found :(")
		}
	}
}