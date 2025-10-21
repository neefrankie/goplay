package main

import (
	"gopl/ch5"
	"os"
)

func main() {
	ch5.BreadthFirst(ch5.Crawl, os.Args[1:])
}
