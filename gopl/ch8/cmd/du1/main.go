package main

import (
	"flag"
	"gopl/ch8/du"
)

func main() {
	flag.Parse()
	roots := flag.Args()

	du.Du1(roots)
}
