// Dup1 prints the text of each line that appears more than
// once in the standard input, preceded by its count.
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	// The scanner reads from the program's standard input.
	input := bufio.NewScanner(os.Stdin)
	// Each call to input.Scan() reads the next line and
	// removes the newline character from the end.
	// The Scan function returns true if there is a line
	// and false when there is no more input.
	for input.Scan() {
		counts[input.Text()]++
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
