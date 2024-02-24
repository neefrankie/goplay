package main

import (
	"log"
	"net"
	"os"

	"example.com/gopl/concurrency"
)

// clock3 accepts a port number from command line argument.
func main() {
	if len(os.Args) < 2 {
		log.Fatal("go run . [PORT]")
	}

	port := os.Args[1]
	listener, err := net.Listen("tcp", "localhost:"+port)
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}

		go concurrency.SendTime(conn)
	}
}
