package main

import (
	"flag"
	"gopl/ch8/du"
)

var verbose = flag.Bool("v", false, "show verbose progress messages")

func main() {
	flag.Parse()
	roots := flag.Args()

	du.Du3(roots, *verbose)
}
