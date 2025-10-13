package netcat

import (
	"log"
	"net"
	"os"
)

func Netcat(addr string) {
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()

	// Reads the standard input and sends it to the server.
	MustCopy(os.Stdout, conn)
}

func Netcat2(addr string) {
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()
	// 把服务器的输出复制到标准输出
	go MustCopy(os.Stdout, conn)
	// 把标准输入复制到服务端
	MustCopy(conn, os.Stdin)
}
