package main

import (
	"io"
	"log"
	"net"
	"os"

	"gopl/concurrency"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	done := make(chan struct{})

	go func() {
		// When the connection closed, a read from closed connection error occur.
		io.Copy(os.Stdout, conn)
		log.Println("done")
		// The main goroutine wait until it has received this value.
		done <- struct{}{}
	}()

	// When user closes the standard input stream, MustCopy returns
	concurrency.MustCopy(conn, os.Stdin)
	// conn.Close()
	if c, ok := conn.(*net.TCPConn); ok {
		// To input end of file, press Ctrl-D on Mac/Linux
		// Ctrl-Z + Return on Windows
		c.CloseWrite()
	} else {
		// then the main gorouine calls the following, the server will see
		// and end-of-file condition
		conn.Close()
	}
	<-done
}
