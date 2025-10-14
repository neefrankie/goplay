package links

import (
	"fmt"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

func Findlinks1() {
	url := os.Args[1]

	doc := crawlAndParse(url)

	for _, link := range visit(nil, doc) {
		fmt.Println(link)
	}
}

func FindLinksRecur() {
	url := os.Args[1]

	doc := crawlAndParse(url)

	for _, link := range visitRecur(nil, doc) {
		fmt.Println(link)
	}
}

func crawlAndParse(url string) *html.Node {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(1)
	}

	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %vn", err)
		os.Exit(1)
	}

	return doc
}
