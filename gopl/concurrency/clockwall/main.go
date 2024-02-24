package main

import (
	"errors"
	"io"
	"log"
	"net"
	"os"
	"strings"
)

func main() {
	args := os.Args[1:]
	for _, arg := range args {
		go start(arg)
	}
}

func start(target string) {
	ent, err := parseEntry(target)
	if err != nil {
		log.Fatal(err)
	}

	catClient(ent)
}

type entry struct {
	loc  string
	addr string
	port string
}

func parseEntry(e string) (entry, error) {
	parts := strings.Split(e, "=")
	if len(parts) < 2 {
		return entry{}, errors.New("format should be locaotion=addr:port")
	}

	address := strings.Split(parts[1], ":")
	if len(address) < 2 {
		return entry{}, errors.New("address must contain a port")
	}

	return entry{
		loc:  parts[0],
		addr: parts[1],
		port: address[1],
	}, nil
}

func catClient(ent entry) {
	conn, err := net.Dial("tcp", ent.addr)
	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()

	mustCopy(os.Stdout, conn)
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
