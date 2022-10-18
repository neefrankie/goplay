package main

import (
	"fmt"
	"goplay/gopl/pkg"
	"os"
)

func main() {
	for _, url := range os.Args[1:] {
		b, err := pkg.Fetch(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			os.Exit(1)
		}
		fmt.Printf("%s\n", b)
	}
}
