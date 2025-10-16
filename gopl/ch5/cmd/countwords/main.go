package main

import (
	"fmt"
	"gopl/ch5"
	"os"
)

func main() {
	for _, url := range os.Args[1:] {
		words, images, err := ch5.CountWordsAndImages(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "countWordsAndImages: %v\n", err)
		}

		fmt.Printf("%s: %d words, %d images\n", url, words, images)
	}
}
