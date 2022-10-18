package html

type Meta struct {
	token Token
	attr  Attr
}

func NewMeta() *Meta {
	return &Meta{
		token: NewToken("meta"),
		attr:  Attr{},
	}
}

func (m *Meta) SetAttr(attr Attr) *Meta {
	m.attr = attr

	return m
}

func (m *Meta) SetHttpEquiv(name, value string) *Meta {
	m.attr = NewAttr().
		Set("http-equiv", name).
		Set("content", value)

	return m
}

func (m *Meta) UACompatible(value string) *Meta {
	if value == "" {
		value = "ie=edge"
	}

	m.SetHttpEquiv("X-UA-Compatible", value)

	return m
}

func (m *Meta) ViewPort(value string) *Meta {
	if value == "" {
		value = "width=device-width, initial-scale=1.0"
	}

	m.attr = NewAttr().
		Set("name", "viewport").
		Set("content", value)

	return m
}

func (m *Meta) CharSet() *Meta {
	m.attr = NewAttr().
		Set("charset", "utf-8")

	return m
}

func (m *Meta) TokenLiteral() string {
	return m.token.Literal
}

func (m *Meta) String() string {
	return NewTag(m.token).
		SelfClosing().
		WithAttr(m.attr).
		String()
}
