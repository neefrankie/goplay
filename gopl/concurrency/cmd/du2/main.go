package main

import (
	"flag"
	"gopl/concurrency/du"
)

var verbose = flag.Bool("v", false, "show verbose progress messages")

func main() {
	flag.Parse()
	roots := flag.Args()

	du.Du2(roots, *verbose)
}
