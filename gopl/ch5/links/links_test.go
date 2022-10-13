package links

import (
	"fmt"
	"golang.org/x/net/html"
	"io"
	"testing"
)

var depth int

func startElement(n *html.Node) {
	if n.Type == html.ElementNode {
		fmt.Printf("%*s<%s>\n", depth*2, "", n.Data)
		depth++
	}
}

func endElement(n *html.Node) {
	if n.Type == html.ElementNode {
		depth--
		fmt.Printf("%*s</%s>\n", depth*2, "", n.Data)
	}
}

func parseHtml(b io.ReadCloser) *html.Node {
	doc, err := html.Parse(b)
	b.Close()
	if err != nil {
		panic(err)
	}

	return doc
}

func getRustHome() io.ReadCloser {

	resp, err := fetch("https://www.rust-lang.org/")
	if err != nil {
		panic(err)
	}

	return resp.Body
}

func Test_forEachNode(t *testing.T) {
	doc := parseHtml(getRustHome())
	forEachNode(doc, startElement, endElement)
}
