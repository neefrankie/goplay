package main

import (
	"fmt"
	"gopl/ch5"
)

func main() {
	for i, course := range ch5.TopoSort1(ch5.Prereqs) {
		fmt.Printf("%d: %s\n", i+1, course)
	}
}
