package html

type Para struct {
	token   Token
	content []Node
}

func NewPara(content []Node) *Para {
	return &Para{
		token:   NewToken("p"),
		content: content,
	}
}

func (p *Para) TokenLiteral() string {
	return p.token.Literal
}

func (p *Para) String() string {
	return NewTag(p.token).WithChildren(p.content).String()
}
