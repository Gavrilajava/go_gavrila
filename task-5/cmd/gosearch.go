package main

import (
	"flag"
	"fmt"
	"go-gavrila/task-5/pkg/crawler/spider"
	"go-gavrila/task-5/pkg/index"
	"log"
	"strings"
)

const depth = 2

var urls = [2]string{"https://go.dev", "https://golang.org"}

func main() {

	token := flag.String("s", "", "search string")
	flag.Parse()

	index, err := index.New()
	if err != nil {
		fmt.Println(err)
	} else {

		if index.Empty() {
			s := spider.New()
			for _, url := range urls {
				res, err := s.Scan(url, depth)
				if err != nil {
					log.Println(err)
					continue
				}
				index = index.Add(res)
			}
			if err = index.Save(); err != nil {
				log.Println(err)
			}
		}

		col := index.Collect(strings.ToLower(*token))

		if len(col) > 0 {
			fmt.Println("Search Results:")
			for _, d := range index.Collect(strings.ToLower(*token)) {
				fmt.Println(d.URL, d.Title)
			}
		} else {
			fmt.Println("Nothing found :(")
		}

	}

}
