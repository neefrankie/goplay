package html

type Token struct {
	Literal string
}

func NewToken(name string) Token {
	return Token{name}
}
