package main

import (
	"flag"
	"gopl/ch8/netcat"
)

func main() {
	flag.Parse()

	args := flag.Args()

	var addr string
	if len(args) == 0 {
		addr = "localhost:8000"
	} else {
		addr = args[0]
	}

	netcat.Netcat(addr)
}
