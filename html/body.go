package html

type Body struct {
	token   Token
	content []Node
}

func NewBody() *Body {
	return &Body{
		token:   NewToken("body"),
		content: []Node{},
	}
}

func (b *Body) WithContent(nodes []Node) *Body {
	b.content = nodes

	return b
}

func (b *Body) String() string {
	return b.token.Literal
}

func (b *Body) TokenLiteral() string {
	return NewTag(b.token).WithChildren(b.content).String()
}
