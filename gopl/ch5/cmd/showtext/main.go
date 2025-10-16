package main

import (
	"gopl/ch5"
	"os"
)

func main() {
	url := os.Args[1]

	doc := ch5.MustCrawlAndParse(url)

	ch5.ShowText(doc)
}
