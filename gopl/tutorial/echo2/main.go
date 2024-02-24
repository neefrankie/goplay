package main

import (
	"fmt"
	"os"
)

// The `for` loop is the only loop statement in Go.
//
// for initialization; condition; post {

// }
//
// Parenthenses are never used around the three components of a for loop.
// The opitonal initialization statement is exeduted before the loop starts.
// If it is present, it must be a simple statement.
// The condition is a boolean expresion that is evaluated at the beginning of each iteration
// of the loop; if it evaluates to true, the statements controlled by the loop are evaluated again.
// The loops ends when the  condition becomes false.
//
// Any of these parts may be omitted. If there is no initialization and no post,
// the semicolon may be omitted:
// for condition {
//
// }
//
// If the condition is omitted entirely in any of the these forms:
//
// for {
//
// }
//
// The loop is inifinite.
// Another form of the loop iterates over a range of values from a data type
// like a string or a slice.
func main() {
	s, sep := "", ""
	// In each iteration of the loop, `range` produces a pair of values:
	// the index and the value of element at that index.
	// The syntax of a `range` loop requires that if we deal with the element,
	// we must deal with the index too.
	// The blank identifier may be used whenever syntax requires a variable name but
	// program logic does not.
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}
