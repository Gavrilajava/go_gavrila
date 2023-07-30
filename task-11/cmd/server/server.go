package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"time"

	"go-gavrila/task-11/pkg/crawler/spider"
	"go-gavrila/task-11/pkg/index"
)

const depth = 2

var urls = [2]string{"https://go.dev", "https://golang.org"}

func handler(conn net.Conn, index index.Service) {
	defer conn.Close()
	defer fmt.Println("Connection Closed")

	conn.SetDeadline(time.Now().Add(time.Second * 30))

	r := bufio.NewReader(conn)

	msg, _, err := r.ReadLine()
	if err != nil {
		return
	}

	for _, d := range index.Collect(strings.ToLower(string(msg))) {
		_, err = conn.Write([]byte(fmt.Sprintf("%s %s \n", d.URL, d.Title)))
		if err != nil {
			return
		}
	}

}

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

	listener, err := net.Listen("tcp4", "0.0.0.0:3500")
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()

	fmt.Println("Server listens on port 3500")

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go handler(conn, *index)
	}
}
