package clock

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

func FTP() {
	listener, err := net.Listen("tcp", ":21")
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

func handleConn(conn net.Conn) {
	defer conn.Close()
	fmt.Fprintf(conn, "220 Welcome to My FTP Server\r\n")

	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		cmd := strings.TrimSpace(scanner.Text())
		if cmd == "" {
			continue
		}

		parts := strings.Split(cmd, " ")
		switch strings.ToUpper(parts[0]) {
		case "USER":
			fmt.Fprintf(conn, "331 user name okay, need password\r]n")
		case "PASS":
			fmt.Fprintf(conn, "230 User logged in\r\n")
		case "QUIT":
			fmt.Fprintf(conn, "221 Goodbye\r\n")
			return
		case "PWD":
			fmt.Fprintf(conn, "257 \"/\" \r\n")
		case "LIST":
			files, _ := os.ReadDir(".")
			for _, f := range files {
				fmt.Fprintf(conn, "-rwxr-xr-x 1 user group 1024 Jan 1 00:00 %s\r\n", f.Name())
			}
			fmt.Fprint(conn, "226 Directory send OK\r\n")
		default:
			fmt.Fprintf(conn, "500 Unknown command\r\n")
		}
	}
}
