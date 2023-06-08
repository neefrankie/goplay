package html

type Text struct {
	Text string
}

func (t *Text) TokenLiteral() string {
	return ""
}

func (t *Text) String() string {
	return t.Text
}
