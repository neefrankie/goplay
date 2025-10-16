package ch5

import (
	"fmt"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

func Findlinks1() {
	url := os.Args[1]

	doc := MustCrawlAndParse(url)

	for _, link := range visit(nil, doc) {
		fmt.Println(link)
	}
}

func FindLinksRecur() {
	url := os.Args[1]

	doc := MustCrawlAndParse(url)

	for _, link := range visitRecur(nil, doc) {
		fmt.Println(link)
	}
}

func FindLinksAll() {
	url := os.Args[1]

	doc := MustCrawlAndParse(url)

	for _, link := range visit2(nil, doc) {
		fmt.Println(link)
	}
}

func FindLinks2(url string) ([]string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("getting %s: %s", url, resp.Status)
	}

	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		return nil, fmt.Errorf("parsing %s as HTML: %v", url, err)
	}

	return visit(nil, doc), nil
}

func CountWordsAndImages(url string) (words, images int, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return
	}

	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		err = fmt.Errorf("parsing HTML: %s", err)
		return
	}

	words, images = countWordsAndImages(doc)
	return
}
