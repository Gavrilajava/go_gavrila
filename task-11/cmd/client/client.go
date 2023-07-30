package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"
)

func main() {

	token := flag.String("s", "", "search string")
	flag.Parse()

	conn, err := net.Dial("tcp4", "0.0.0.0:3500")
	if err != nil {
		log.Fatal(err)
	}

	_, err = conn.Write(append([]byte(*token), '\n'))
	if err != nil {
		log.Fatal(err)
	}

	msg, err := io.ReadAll(conn)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Server response: \n", string(msg))
}
