package lib

import (
	"net/http"

	"golang.org/x/net/html"
)

func CrawlAndParse(url string) (*html.Node, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return nil, err
	}

	return doc, nil
}

func MustCrawlAndParse(url string) *html.Node {
	doc, err := CrawlAndParse(url)
	if err != nil {
		panic(err)
	}

	return doc
}
