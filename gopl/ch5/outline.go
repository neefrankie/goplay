package ch5

import (
	"fmt"
	"golang.org/x/net/html"
	"io"
)

func outline(stack []string, n *html.Node) {
	if n.Type == html.ElementNode {
		stack = append(stack, n.Data)
		fmt.Println(stack)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		outline(stack, c)
	}
}

func htmlOutline(in io.Reader) error {
	doc, err := html.Parse(in)
	if err != nil {
		return err
	}
	outline(nil, doc)
	return nil
}
