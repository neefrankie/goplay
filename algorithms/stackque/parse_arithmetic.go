package stackque

import (
	"strings"
)

const (
	_ int = iota
	PrecedenceSum
	PrecedenceProduct
)

// Operator precedence
const (
	_ int = iota
	sum
	product
)

var precedences = map[byte]int{
	'+': sum,
	'-': sum,
	'*': product,
	'/': product,
}

// PostfixBuilder converts infix arithmetic expression to postfix form.
type PostfixBuilder struct {
	operators *Stack[byte]
	input     string
	output    strings.Builder
}

func NewPostfixBuilder(input string) *PostfixBuilder {
	return &PostfixBuilder{
		operators: NewStack[byte](len(input)),
		input:     input,
		output:    strings.Builder{},
	}
}

func (b *PostfixBuilder) Convert() string {
	length := len(b.input)

	for i := 0; i < length; i++ {

		if b.input[i] >= '0' && b.input[i] <= '9' {
			if i != 0 {
				b.output.WriteByte(' ')
			}

			for i < length && b.input[i] >= '0' && b.input[i] <= '9' {
				b.output.WriteByte(b.input[i])
				i++
			}
			// Go one step backward so that the outer loop i++
			// won't jump over next char
			i--
		} else {
			ch := b.input[i]

			switch ch {
			case '+', '-', '*', '/':
				b.pushOperator(ch, precedences[ch])

			case '(':
				_ = b.operators.Push(ch)

			case ')':
				for !b.operators.IsEmpty() {
					chx, _ := b.operators.Pop()
					// If popped (, we're done.
					if chx == '(' {
						break
					} else {
						// Otherwise output operators.
						b.output.WriteByte(chx)
					}
				}

			default:

			}
		}
	}

	for !b.operators.IsEmpty() {
		ch, _ := b.operators.Pop()
		b.output.WriteByte(' ')
		b.output.WriteByte(ch)
	}

	return b.output.String()
}

func (b *PostfixBuilder) pushOperator(newOp byte, precedence int) {
	for !b.operators.IsEmpty() {
		opTop, _ := b.operators.Pop()
		// If it's a (, restore it.
		if opTop == '(' {
			_ = b.operators.Push(opTop)
			break
		} else {
			var precTop = precedences[opTop]
			// If new operator has a higher precedence
			if precedence > precTop {
				_ = b.operators.Push(opTop)
				break
			} else {
				// Add the popped operator to output.
				b.output.WriteByte(' ')
				b.output.WriteByte(opTop)
			}
		}
	}

	_ = b.operators.Push(newOp)
}

func readNum(input string, pos int) int {
	for input[pos] >= '0' && input[pos] <= '9' {
		pos++
	}

	return pos
}

func evaluatePostfix(input string) int {
	stack := NewStack[int](len(input))
	l := len(input)

	for i := 0; i < l; i++ {
		if input[i] >= '0' && input[i] <= '9' {
			_ = stack.Push(0)
			for ch := input[i]; ch >= '0' && ch <= '9' && i < l; i++ {
				n, _ := stack.Pop()
				_ = stack.Push(10*n + int(input[i]-'0'))
			}
		} else if isOperator(input[i]) {
			var result int

			num2, _ := stack.Pop()
			num1, _ := stack.Pop()
			switch input[i] {
			case '+':
				result = num1 + num2
			case '-':
				result = num1 - num2
			case '*':
				result = num1 * num2
			case '/':
				result = num1 / num2
			}
			_ = stack.Push(result)
		}
	}

	result, _ := stack.Pop()
	return result
}

func isOperator(ch byte) bool {
	return ch == '+' || ch == '-' || ch == '*' || ch == '/'
}
