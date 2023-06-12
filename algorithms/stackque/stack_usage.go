package stackque

import (
	"errors"
	"fmt"
	"strings"
)

// ReverseWord reverse a word using stack
func ReverseWord(word string) string {
	stack := NewStack[rune](len(word))
	for _, r := range word {
		_ = stack.Push(r)
	}

	b := strings.Builder{}
	for !stack.IsEmpty() {
		r, _ := stack.Pop()
		b.WriteRune(r)
	}

	return b.String()
}

type BracketError struct {
	Char     rune
	Position int
}

func (e *BracketError) Error() string {
	return fmt.Sprintf("bracket checker error: %s at %d", string(e.Char), e.Position)
}

// CheckBracket uses a stack to parse a string and checks if there's any mismatched brackets.
func CheckBracket(input string, handler func(err *BracketError)) error {
	stack := NewStack[rune](len(input))

	for i, ch := range input {
		switch ch {
		case '{', '[', '(':
			_ = stack.Push(ch)

		case '}', ']', ')':
			if !stack.IsEmpty() {
				chx, _ := stack.Pop()
				if (ch == '}' && chx != '{') || (ch == ']' && chx != '[') || (ch == ')' && chx != '(') {
					handler(&BracketError{
						Char:     ch,
						Position: i,
					})
				}
			} else {
				handler(&BracketError{
					Char:     ch,
					Position: i,
				})
			}

		default:

		}
	}

	if !stack.IsEmpty() {
		return errors.New("missing right delimiter")
	}

	return nil
}
