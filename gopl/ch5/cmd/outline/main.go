package main

import (
	"gopl/ch5/lib"
	"gopl/ch5/outline"
)

func main() {
	doc := lib.MustCrawlAndParse("http://127.0.0.1:5000")

	outline.Outline(nil, doc)
}
