package main

import (
	"fmt"
	"gopl/ch5"
	"os"
)

func main() {
	for _, url := range os.Args[1:] {
		links, err := ch5.FindLinks2(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "findlinks2: %v\n", err)
		}

		for _, link := range links {
			fmt.Println(link)
		}
	}
}
