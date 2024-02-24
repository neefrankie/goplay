package concurrency

import (
	"fmt"
	"io"
	"net"
	"strings"
	"time"
)

func SendTime(c net.Conn) {
	defer c.Close()
	// In the loop, it writes teh current time to client.
	// Since `net.Conn` satisfies the `io.Writer` interface,
	//  we can write directly to it.
	for {
		// The loop ends when the write fails, most likely because the client has disconnected,
		// at which point `handleConn` closes its side of the connection using a deferred call
		// to `Close` and goes back to waiting for another connection request.
		_, err := io.WriteString(c, time.Now().Format("15:04:05\n"))
		if err != nil {
			return
		}
		time.Sleep(1 * time.Second)
	}
}

func Echo(c net.Conn, shout string, delay time.Duration) {
	fmt.Fprintln(c, "\t", strings.ToUpper(shout))
	time.Sleep(delay)

	fmt.Fprintln(c, "\t", shout)
	time.Sleep(delay)

	fmt.Fprintln(c, "\t", strings.ToLower(shout))
}
