package main

import (
	"fmt"
	"os"
	"playground/gopl/ch1"
)

func main() {
	for _, url := range os.Args[1:] {
		b, err := ch1.Fetch(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			os.Exit(1)
		}
		fmt.Printf("%s\n", b)
	}
}
