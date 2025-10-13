package clock

import (
	"io"
	"log"
	"net"
	"time"
)

// A sequential clock server that writes the current time to the client once persecond.
func SequentialClock(port string) {
	// `Listen` creates a `net.Listener`, an object that listens for incoming connections on a network port.
	listener, err := net.Listen("tcp", "localhost:"+port)
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
		HandleConn(conn)
	}
}

func ParalellClock(port string) {
	// `Listen` creates a `net.Listener`, an object that listens for incoming connections on a network port.
	listener, err := net.Listen("tcp", "localhost:"+port)
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

		// Adding the go keyword to the call to `handleConn` causes each call
		// to run in its own goroutine.
		go HandleConn(conn)
	}
}

// HandleConn handles on complete client connection.
// Adding the go keyword to the call to `handleConn` causes each call
// to run in its own goroutine.
func HandleConn(c net.Conn) {
	defer c.Close()
	for {
		_, err := io.WriteString(c, time.Now().Format("15:04:05\n"))
		if err != nil {
			return
		}

		time.Sleep(1 * time.Second)
	}
}
