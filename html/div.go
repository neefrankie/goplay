package html

type Div struct {
	token    Token
	children []Node
}

func NewDiv(content []Node) *Div {
	return &Div{
		token:    NewToken("div"),
		children: content,
	}
}

func (d *Div) TokenLiteral() string {
	return d.token.Literal
}

func (d *Div) String() string {
	return NewTag(d.token).WithChildren(d.children).String()
}
