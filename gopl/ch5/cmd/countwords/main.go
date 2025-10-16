package main

import (
	"fmt"
	"gopl/ch5/links"
	"os"
)

func main() {
	for _, url := range os.Args[1:] {
		words, images, err := links.CountWordsAndImages(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "countWordsAndImages: %v\n", err)
		}

		fmt.Printf("%s: %d words, %d images\n", url, words, images)
	}
}
