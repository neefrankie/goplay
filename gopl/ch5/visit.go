package ch5

import (
	"strings"

	"golang.org/x/net/html"
)

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

func visit2(links []string, n *html.Node) []string {
	if n == nil {
		return links
	}

	if n.Type == html.ElementNode {
		switch strings.ToLower(n.Data) {
		case "a", "link":
			for _, a := range n.Attr {
				if a.Key == "href" {
					links = append(links, a.Val)
				}
			}
		case "image", "script":
			for _, a := range n.Attr {
				if a.Key == "src" {
					links = append(links, a.Val)
				}
			}
		}
	}

	links = visit2(links, n.FirstChild)
	links = visit2(links, n.NextSibling)

	return links
}

func visitRecur(links []string, n *html.Node) []string {
	if n == nil {
		return links
	}

	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}

	links = visitRecur(links, n.FirstChild)
	// Each recursive call goes to the next sibling
	links = visitRecur(links, n.NextSibling)

	return links
}
