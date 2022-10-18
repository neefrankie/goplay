package html

import "strings"

type HTML struct {
	token Token
	lang  string
	head  *Head
	body  *Body
}

func NewHTML(lang string) *HTML {
	return &HTML{
		token: NewToken("html"),
		lang:  lang,
		head:  nil,
		body:  nil,
	}
}

func (h *HTML) WithHead(head *Head) *HTML {
	h.head = head

	return h
}

func (h *HTML) WithBody(body *Body) *HTML {
	h.body = body

	return h
}

func (h *HTML) TokenLiteral() string {
	return h.token.Literal
}

func (h *HTML) String() string {
	var out strings.Builder
	out.WriteString("<!DOCTYPE html>")

	tag := NewTag(h.token).
		WithAttr(
			NewAttr().Set("lang", h.lang),
		).
		WithChildren(
			[]Node{
				h.head,
				h.body,
			},
		)

	out.WriteString(tag.String())

	return out.String()
}
