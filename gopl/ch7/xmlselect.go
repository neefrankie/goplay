package ch7

import (
	"encoding/xml"
	"fmt"
	"io"
	"strings"
)

// XmlSelect prints the text of selected elements of an XML document.
func XmlSelect(r io.Reader, selector []string, visitor func(data string)) error {
	dec := xml.NewDecoder(r)
	var stack []string

	// Each time the loop encounters a StartElement, it pushes the element's name
	// onto a stack, and for each EndElement it pops the name from the stack.
	// When it encounters a CharData, it prints the text only if the stack contains
	// all the elements named by selector, in order
	for {
		tok, err := dec.Token()
		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}

		switch tok := tok.(type) {
		case xml.StartElement:
			stack = append(stack, tok.Name.Local)
		case xml.EndElement:
			stack = stack[:len(stack)-1]
		case xml.CharData:
			if containsAll(stack, selector) {
				visitor(fmt.Sprintf("%s: %s\n", strings.Join(stack, " "), tok))
			}
		}
	}

	return nil
}

// containsAll reports whether x contains the elements of y, in order.
func containsAll(x, y []string) bool {
	for len(y) <= len(x) {
		if len(y) == 0 {
			return true
		}
		if x[0] == y[0] {
			y = y[1:]
		}
		x = x[1:]
	}
	return false
}
