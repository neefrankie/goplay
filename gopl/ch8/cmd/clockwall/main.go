package main

import (
	"flag"
	"gopl/ch8/clock"
)

func main() {
	flag.Parse()

	addresses := flag.Args()

	clock.ClockWall(addresses)
}
