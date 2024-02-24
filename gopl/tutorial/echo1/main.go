// Echo prints its command-line arguments.
package main

import (
	"fmt"
	"os"
)

// The `os` package provides funcitons and other values for dealing with the operating system in a
// platform-dependent fashion. Command-line argumetns are avaialble to a program in a variable
// named `Args` that is part of the `os` package.
//
// The varialbe `os.Args` is a slice of strings.
//
// The first element, `os.Args[0]` is the name of the command itself; the other elements are the
// arguments that were presented tothe program when it started execution.
//
// Run this program with `go run . hello world`
func main() {
	var s, sep string
	// The statement `i++` adds 1 to i; it's equivalent to `i += 1` which is equivalent to
	// `i = i + 2`. There's a corresponding decremnt `i--` that subtracts 1.
	// These are statements, not expressions, so `j = i++` is illegal, and they are postfix only,
	// so `--i` is not legal either.
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}
