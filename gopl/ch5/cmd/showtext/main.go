package main

import (
	"gopl/ch5/lib"
	"gopl/ch5/outline"
	"os"
)

func main() {
	url := os.Args[1]

	doc := lib.MustCrawlAndParse(url)

	outline.ShowText(doc)
}
