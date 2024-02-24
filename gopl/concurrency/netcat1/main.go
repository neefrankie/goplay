package main

import (
	"io"
	"log"
	"net"
	"os"
)

// Netcat1 is a read-only TCP client.
// Thie program reads data from the connection and writes it to the standared output
// until an end-of-file condition or an error occurs.
func main() {
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()
	// Reads the standard input and sends it to the server.
	mustCopy(os.Stdout, conn)
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
