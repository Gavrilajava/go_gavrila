package netsrv

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"time"

	"go-gavrila/task-11/pkg/index"
)

type Server struct {
	port  string
	index index.Service
}

func New(p string, index index.Service) *Server {
	return &Server{p, index}
}

func (s Server) Start() {
	listener, err := net.Listen("tcp4", "0.0.0.0:"+s.port)
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()

	fmt.Println("Server listens on port 8000")

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go s.handler(conn)
	}
}

func (s Server) handler(conn net.Conn) {
	defer conn.Close()
	defer fmt.Println("Connection Closed")

	r := bufio.NewReader(conn)

	for {
		msg, _, err := r.ReadLine()
		if err != nil {
			return
		}
		if len(msg) == 0 {
			break
		}

		fmt.Println("Searching:", string(msg))

		result := s.index.Collect(strings.ToLower(string(msg)))

		if len(result) == 0 {
			if _, err := fmt.Fprintf(conn, "Nothing found\n\n"); err != nil {
				return
			}
		}

		for _, d := range result {
			if _, err := fmt.Fprintf(conn, "%s %s\n", d.URL, d.Title); err != nil {
				return
			}
		}
		if _, err := fmt.Fprintf(conn, "\n"); err != nil {
			return
		}

		conn.SetDeadline(time.Now().Add(time.Second * 30))
	}

}
