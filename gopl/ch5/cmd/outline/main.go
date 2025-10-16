package main

import (
	"gopl/ch5"
)

func main() {
	doc := ch5.MustCrawlAndParse("http://127.0.0.1:5000")

	ch5.Outline(nil, doc)
}
