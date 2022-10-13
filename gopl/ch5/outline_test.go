package ch5

import (
	"bytes"
	"golang.org/x/net/html"
	"testing"
)

func parseHtml(b []byte) *html.Node {
	doc, err := html.Parse(bytes.NewReader(b))
	if err != nil {
		panic(err)
	}

	return doc
}

func Test_htmlOutline(t *testing.T) {
	doc := parseHtml(getRustHome())
	outline(nil, doc)
}
