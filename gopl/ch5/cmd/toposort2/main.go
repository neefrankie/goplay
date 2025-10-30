package main

import (
	"fmt"
	"gopl/ch5"
)

func main() {
	for i, course := range ch5.TopoSort2(ch5.PrereqsCycle) {
		fmt.Printf("%d: %s\n", i+1, course)
	}
}
