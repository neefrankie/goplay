package ch5

import (
	"fmt"

	"golang.org/x/net/html"
)

func formatAttr(attr []html.Attribute) string {
	var s, sep string
	for _, a := range attr {
		sep = " "
		s += sep + fmt.Sprintf("%s=%q", a.Key, a.Val)
	}

	return s
}

func FormatOpeningElement(n *html.Node, depth int) string {
	return fmt.Sprintf("%*s<%s%s>\n", depth*2, "", n.Data, formatAttr(n.Attr))
}

func FormatClosingElement(n *html.Node, depth int) string {
	return fmt.Sprintf("%*s</%s>\n", depth*2, "", n.Data)
}

func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}

	if n.FirstChild != nil {
		if post != nil {
			post(n)
		}
	}
}

func Outline2(doc *html.Node) {
	var depth int

	var startElement = func(n *html.Node) {
		if n.Type == html.ElementNode {
			// %*s 中的 * 会在字符串之前填充一些空格
			if n.FirstChild != nil {
				// fmt.Printf("%*s<%s%s>\n", depth*2, "", n.Data, formatAttr(n.Attr))
				print(FormatOpeningElement(n, depth))
				depth++
			} else {
				fmt.Printf("%*s<%s%s/>\n", depth*2, "", n.Data, formatAttr(n.Attr))
			}
		}
	}

	var endElement = func(n *html.Node) {
		if n.Type == html.ElementNode {
			depth--
			// fmt.Printf("%*s</%s>\n", depth*2, "", n.Data)
			print(FormatClosingElement(n, depth))
		}
	}

	forEachNode(doc, startElement, endElement)
}
