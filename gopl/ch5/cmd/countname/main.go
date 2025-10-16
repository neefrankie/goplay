package main

import (
	"fmt"
	"gopl/ch5/lib"
	"gopl/ch5/outline"
	"os"
)

func main() {
	url := os.Args[1]

	doc := lib.MustCrawlAndParse(url)

	counter := make(map[string]int)

	outline.CountNodeNames(counter, doc)

	for name, count := range counter {
		fmt.Printf("%s: %d\n", name, count)
	}
}
