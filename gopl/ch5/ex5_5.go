package ch5

import (
	"regexp"
	"strings"

	"golang.org/x/net/html"
)

func forEachNodeFind(n *html.Node, f func(n *html.Node) bool) *html.Node {
	if f != nil {
		ok := f(n)
		if !ok {
			return n
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if n := forEachNodeFind(c, f); n != nil {
			return n
		}
	}

	return nil
}

func ElementByID(doc *html.Node, id string) *html.Node {
	var f = func(n *html.Node) bool {
		if n.Type == html.ElementNode {
			for _, a := range n.Attr {
				if a.Key == "id" {
					return false
				}
			}
		}
		return true
	}

	return forEachNodeFind(doc, f)
}

func expand(s string, f func(string) string) string {
	regex, err := regexp.Compile(`foo`)
	if err != nil {
		panic(err)
	}

	s = regex.ReplaceAllStringFunc(s, f)

	return s
}

func expand2(s string, f func(string) string) string {
	parts := strings.Split(s, "foo")

	return strings.Join(parts, f("foo"))
}

func exapnd3(s string, f func(string) string) string {
	return strings.ReplaceAll(s, "foo", f("foo"))
}
