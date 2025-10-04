package main

import (
	"flag"
	"fmt"

	"gopl/ch7"
)

var temp = ch7.CelsiusFlag("temp", 20.0, "the temperature")

func main() {
	flag.Parse()
	fmt.Println(*temp)
}
