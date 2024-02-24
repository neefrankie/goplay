package main

import (
	"bufio"
	"fmt"
	"os"
)

// Use bufio.Scanner to read standard input as a set of lines.
func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		// Println will add back the final `\n`
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input: ", err)
	}
}
