package main

import (
	"gopl/ch5"
	"os"
)

func main() {

	url := os.Args[1]

	doc := ch5.MustCrawlAndParse(url)
	elem := ch5.ElementByID(doc, "top")
	if elem != nil {
		print(ch5.FormatOpeningElement(elem, 0))
	} else {
		print("No element with id 'top' found")
	}
}
