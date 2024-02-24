package main

import (
	"log"
	"net"
	"os"

	"example.com/gopl/concurrency"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()

	// Copy the server response to the ouput.
	go concurrency.MustCopy(os.Stdout, conn)

	// Send terminal input to the server.
	// While the main goroutine reads the standard input and sends it to the server,
	// a second goroutine reads and prints the server's response. When the main
	// goroutine encounters the end of the input, the program stops, even if the other
	// goroutine still has work to do.
	concurrency.MustCopy(conn, os.Stdin)
}
