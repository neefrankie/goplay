package main

import (
	"fmt"
	"gopl/ch5"
	"os"
)

func main() {
	url := os.Args[1]

	doc := ch5.MustCrawlAndParse(url)

	counter := make(map[string]int)

	ch5.CountNodeNames(counter, doc)

	for name, count := range counter {
		fmt.Printf("%s: %d\n", name, count)
	}
}
