package main

import (
	"flag"
	"fmt"
	"go-gavrila/task-5/pkg/crawler/spider"
	"go-gavrila/task-5/pkg/index"
	"strings"
)

const depth = 2

var urls = [2]string{"https://go.dev", "https://golang.org"}

func main() {

	token := flag.String("s", "", "search string")
	flag.Parse()
	scan()
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

func scan() {
	index.LoadFiles()
	if index.Empty() {
		fmt.Println("Searching...")
		s := spider.New()
		for _, url := range urls {
			res, err := s.Scan(url, depth)
			if err != nil {
				fmt.Println(err)
				return
			} else {
				index.Add(res)
			}
		}
	}
}
