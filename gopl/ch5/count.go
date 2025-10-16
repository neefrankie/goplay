package ch5

import (
	"fmt"
	"strings"
	"unicode"

	"golang.org/x/net/html"
)

// CountNodeNames implements exercise 5.2s
func CountNodeNames(counter map[string]int, n *html.Node) {
	if n.Type == html.ElementNode {
		counter[n.Data]++
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		CountNodeNames(counter, c)
	}
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
