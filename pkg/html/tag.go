package html

import "strings"

type Tag struct {
	token       Token
	selfClosing bool
	attr        Attr
	children    []Node
}

func NewTag(token Token) *Tag {
	return &Tag{
		token:       token,
		selfClosing: false,
		attr:        nil,
		children:    nil,
	}
}

func (t *Tag) SelfClosing() *Tag {
	t.selfClosing = true
	return t
}

func (t *Tag) WithAttr(attr Attr) *Tag {
	t.attr = attr

	return t
}

func (t *Tag) WithChildren(nodes []Node) *Tag {
	t.children = nodes

	return t
}

func (t *Tag) AddChildren(nodes []Node) *Tag {
	for _, v := range nodes {
		t.children = append(t.children, v)
	}

	return t
}

func (t *Tag) AddChild(node Node) *Tag {
	t.children = append(t.children, node)

	return t
}

func (t *Tag) leading() string {
	var out strings.Builder

	out.WriteByte('<')
	out.WriteString(t.token.Literal)
	if !t.attr.IsEmpty() {
		out.WriteByte(' ')
		out.WriteString(t.attr.Encode())
	}

	out.WriteByte('>')

	return out.String()
}

func (t *Tag) trailing() string {
	if t.selfClosing {
		return ""
	}
	return "</" + t.token.Literal + ">"
}

func (t *Tag) TokenLiteral() string {
	return t.token.Literal
}

func (t *Tag) String() string {
	var out strings.Builder

	out.WriteString(t.leading())

	for _, node := range t.children {
		out.WriteString(node.String())
	}

	out.WriteString(t.trailing())

	return out.String()
}
