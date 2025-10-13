package main

import (
	"gopl/ch8/reverb"
	"net"
)

func main() {
	conn, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		panic(err)
	}

	for {
		conn, err := conn.Accept()
		if err != nil {
			panic(err)
		}
		go reverb.HandleReverb1(conn)
	}
}
