package main

import (
	"fmt"
	"log"
	"os"

	"example.com/gopl/func/links"
)

var tokens = make(chan struct{}, 20)

func crawl(url string) []string {
	fmt.Println(url)

	// Acquire a token
	tokens <- struct{}{}
	list, err := links.Extract(url)
	// Release the token
	<-tokens

	if err != nil {
		log.Print(err)
	}

	return list
}

func main() {
	worklist := make(chan []string)
	var n int // number of pending sends to worklist

	// Start with teh command-line arguments
	n++
	go func() {
		worklist <- os.Args[1:]
	}()

	seen := make(map[string]bool)
	for ; n > 0; n-- {
		for list := range worklist {
			for _, link := range list {
				if !seen[link] {
					seen[link] = true
					n++
					go func(link string) {
						worklist <- crawl(link)
					}(link)
				}
			}
		}
	}
}
