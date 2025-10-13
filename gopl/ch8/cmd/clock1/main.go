package main

import (
	"flag"
	"gopl/ch8/clock"
)

var port = flag.String("port", "8000", "-port=8000")

// A sequential clock server that writes the current time to the client once persecond.
// Access this server:
// nc localhost 8000
// Or use the ./netcat.
func main() {
	flag.Parse()

	clock.SequentialClock(*port)
}
