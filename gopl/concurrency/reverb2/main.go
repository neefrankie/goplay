package main

import (
	"bufio"
	"log"
	"net"
	"os"
	"time"

	"example.com/gopl/concurrency"
)

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

		go handleConn(conn)
	}
}

func handleConn(c net.Conn) {
	input := bufio.NewScanner(c)
	for input.Scan() {
		go concurrency.Echo(c, input.Text(), 1*time.Second)
	}

	c.Close()
}
