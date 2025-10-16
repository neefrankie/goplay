package links

import (
	"fmt"
	"strings"
	"unicode"

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

func countWordsAndImages(n *html.Node) (words, images int) {

	if n == nil {
		return
	}

	if n.Type == html.TextNode {
		text := strings.TrimSpace(n.Data)
		if text != "" {
			// w := strings.Fields(text)
			n := countWords(text)
			words += n
			fmt.Printf("%s: %d\n", text, n)
		}
	}

	if n.Type == html.ElementNode {
		lowerData := strings.ToLower(n.Data)
		if lowerData == "img" {
			images++
		}

		if lowerData != "script" && lowerData != "style" {
			w, img := countWordsAndImages(n.FirstChild)
			words += w
			images += img
		}
	} else {
		w, img := countWordsAndImages(n.FirstChild)
		words += w
		images += img
	}

	w, img := countWordsAndImages(n.NextSibling)
	words += w
	images += img

	return
}

// countWords counts the number of words in a string
// 这里的实现是不严谨的。
func countWords(text string) int {
	var count int
	for _, r := range text {
		// 只统计字母和汉字等可读字符
		if unicode.IsLetter(r) || unicode.IsNumber(r) {
			count++
		}
	}
	return count
}
