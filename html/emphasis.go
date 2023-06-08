package html

type Emphasis struct {
	token Token
	text  *Text
}

func NewEmphasis(text string) *Emphasis {
	return &Emphasis{
		token: NewToken("em"),
		text:  &Text{text},
	}
}

func (e *Emphasis) TokenLiteral() string {
	return e.token.Literal
}

func (e *Emphasis) String() string {
	return NewTag(e.token).
		WithChildren(
			[]Node{e.text},
		).
		String()
}
