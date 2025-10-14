package main

import (
	"fmt"
	"gopl/func/links"
	"log"
	"os"
)

// Each of the n vacant slots in the channel buffer represents a token
// entitling the hodler to proceed.
// Sending a value into the channel acquires a token, and receiving a value
// from the channel releases a token, creating a new vacant slot.
var tokens = make(chan struct{}, 20)

func crawl(url string) []string {
	fmt.Println(url)
	// Send a value to acquire a token.
	tokens <- struct{}{}
	list, err := links.Extract(url)
	// Receive a value to release a token.
	<-tokens
	if err != nil {
		log.Print(err)
	}

	return list
}

func main() {
	worklist := make(chan []string)
	var n int

	n++

	go func() {
		worklist <- os.Args[1:]
	}()

	seen := make(map[string]bool)
	for ; n > 0; n-- {
		list := <-worklist
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
