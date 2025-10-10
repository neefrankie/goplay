package main

import (
	"flag"
	"gopl/concurrency/du"
)

func main() {
	flag.Parse()
	roots := flag.Args()

	du.Du1(roots)
}
