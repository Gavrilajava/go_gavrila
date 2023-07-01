package main

import (
	"go-gavrila/task-3/pkg/crawler/spider"
	"go-gavrila/task-3/pkg/index"
	"flag"
	"strings"
	"fmt"
)

const depth = 2
var urls = [2]string{"https://go.dev", "https://golang.org"}

func main() {

	s := spider.New()
	token := flag.String("s", "", "search string")
	

	for _, url := range urls {
		res, err := s.Scan(url, depth)
		if err != nil {
			fmt.Println(err)
			return
		} else {
			index.Add(res)
		}
	}
	
	flag.Parse()
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