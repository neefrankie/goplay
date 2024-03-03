package main

import (
	"log"
	"net"

	"gopl/concurrency"
)

// A sequential clock server that writes the current time to the client once persecond.
// Access this server:
// nc localhost 8000
// Or use the ./netcat.
func main() {
	// `Listen` creates a `net.Listener`, an object that listens for incoming connections on a network port.
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	for {
		// `Accept` method blocks until an incoming connection request is made,
		// then returns a `net.Conn` object representing the connection.
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		// The handleConn function handles on complete client connection.
		// Adding the go keyword to the call to `handleConn` causes each call
		// to run in its own goroutine.
		concurrency.SendTime(conn)
	}
}
