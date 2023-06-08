package fetch

import (
	"net/url"
	"strings"
)

type URLBuilder struct {
	base     string
	paths    []string
	query    url.Values
	rawQuery string
}

func NewURLBuilder(base string) URLBuilder {
	return URLBuilder{
		base:     base,
		paths:    make([]string, 0),
		query:    make(url.Values),
		rawQuery: "",
	}
}

func (b URLBuilder) SetPath(p string) URLBuilder {
	b.paths = []string{p}
	return b
}

func (b URLBuilder) AddPath(p string) URLBuilder {
	b.paths = append(b.paths, p)
	return b
}

func (b URLBuilder) SetRawQuery(q string) URLBuilder {
	b.rawQuery = q
	return b
}

func (b URLBuilder) SetQuery(k, v string) URLBuilder {
	b.query.Set(k, v)
	return b
}

func (b URLBuilder) AddQuery(k, v string) URLBuilder {
	b.query.Add(k, v)
	return b
}

func (b URLBuilder) String() string {
	var buf strings.Builder
	if b.base != "" {
		buf.WriteString(b.base)
	}

	path := strings.Join(b.paths, "/")
	if b.base != "" && path != "" && !strings.HasPrefix(path, "/") {
		buf.WriteByte('/')
	}

	buf.WriteString(path)

	query := b.query.Encode()

	if b.rawQuery != "" || query != "" {
		buf.WriteByte('?')

		buf.WriteString(b.rawQuery)

		if b.rawQuery != "" {
			buf.WriteByte('&')
		}

		buf.WriteString(query)
	}

	return buf.String()
}
