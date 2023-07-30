package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

func main() {

	conn, err := net.Dial("tcp4", "0.0.0.0:8000")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	defer fmt.Println("Connection Closed")

	server := bufio.NewReader(conn)
	console := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("Eneter your search term:")
		msg, _, err := console.ReadLine()
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("Searching", string(msg))
		fmt.Fprintf(conn, "%s\n", string(msg))

		for {
			reply, _, err := server.ReadLine()
			if err != nil {
				log.Fatal(err)
			}
			if len(reply) == 0 {
				break
			}
			fmt.Println(string(reply))
		}
	}

}
