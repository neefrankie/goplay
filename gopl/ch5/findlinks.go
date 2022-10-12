package ch5

import (
	"fmt"
	"golang.org/x/net/html"
	"io"
)

func findLinks(in io.Reader) ([]string, error) {
	doc, err := html.Parse(in)
	if err != nil {
		return nil, fmt.Errorf("findlinks: %v", err)
	}

	return visit(nil, doc), nil
}

func visit(links []string, n *html.Node) []string {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = visit(links, c)
	}

	return links
}
