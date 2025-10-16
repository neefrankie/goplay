package main

import (
	"fmt"
	"gopl/ch5/links"
	"os"
)

func main() {
	for _, url := range os.Args[1:] {
		links, err := links.FindLinks2(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "findlinks2: %v\n", err)
		}

		for _, link := range links {
			fmt.Println(link)
		}
	}
}
