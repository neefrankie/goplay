package html

type Head struct {
	token Token
	base  *Base
	meta  []Node
	title *Title
}

func NewHead() *Head {
	return &Head{
		token: NewToken("head"),
		base:  nil,
		meta:  nil,
	}
}

func (h *Head) WithBasePath(p string) *Head {
	h.base = NewBase().WithPath(p)

	return h
}

func (h *Head) WithMeta(m *Meta) *Head {

	h.meta = append(h.meta, m)

	return h
}

func (h *Head) WithTitle(t string) *Head {
	h.title = NewTitle(t)

	return h
}

func (h *Head) String() string {
	tag := NewTag(h.token).
		WithChildren(
			[]Node{
				h.base,
				h.title,
			},
		).
		AddChildren(h.meta)

	return tag.String()
}

func (h *Head) TokenLiteral() string {
	return h.token.Literal
}

type Base struct {
	token Token
	attr  Attr
}

func NewBase() *Base {
	return &Base{
		token: NewToken("base"),
		attr:  NewAttr(),
	}
}

func (b *Base) WithPath(p string) *Base {
	b.attr.Set("href", p)

	return b
}

func (b *Base) TokenLiteral() string {
	return b.token.Literal
}

func (b *Base) String() string {
	return NewTag(b.token).SelfClosing().WithAttr(b.attr).String()
}

type Title struct {
	token Token
	text  *Text
}

func NewTitle(t string) *Title {
	return &Title{
		token: NewToken("title"),
		text: &Text{
			Text: t,
		},
	}
}

func (t *Title) TokenLiteral() string {
	return t.token.Literal
}

func (t *Title) String() string {
	return NewTag(t.token).
		WithChildren([]Node{t.text}).
		String()
}
