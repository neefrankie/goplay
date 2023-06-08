package html

import "strings"

// Attr represents all the attributes of an HTML tag.
// It maps a string key to a list of values.
type Attr map[string][]string

func NewAttr() Attr {
	return Attr{}
}

func (a Attr) IsEmpty() bool {
	return len(a) == 0
}

// Set sets the name to value. It replaces any existing values.
// To set a boolean attribute, use empty string as value.
func (a Attr) Set(name, value string) Attr {
	value = strings.TrimSpace(value)

	if value == "" {
		a[name] = []string{}
		return a
	}

	a[name] = []string{value}
	return a
}

// Add adds the value to name. It appends to any existing
// values associated with the name
func (a Attr) Add(name, value string) Attr {
	if value == "" {
		return a
	}
	a[name] = append(a[name], value)

	return a
}

// Get gets the first value associated with the given key.
// If there are no values associated with the key, Get returns
// the empty string.
func (a Attr) Get(key string) string {
	if a == nil {
		return ""
	}
	vs := a[key]
	if len(vs) == 0 {
		return ""
	}

	return vs[0]
}

// Del deletes the values associated with name
func (a Attr) Del(name string) Attr {
	delete(a, name)

	return a
}

// Encode turns the values into tag attribute string.
// class="btn btn-primary" aria-label="hello"
func (a Attr) Encode() string {
	if a == nil {
		return ""
	}

	var buf strings.Builder
	for n, v := range a {
		if buf.Len() > 0 {
			buf.WriteByte(' ')
		}
		buf.WriteString(n)
		if len(v) == 0 {
			continue
		}
		buf.WriteByte('=')
		buf.WriteByte('"')
		buf.WriteString(strings.Join(v, " "))
		buf.WriteByte('"')
	}

	return buf.String()
}
