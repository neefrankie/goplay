package fetch

import (
	"net/url"
	"testing"
)

func TestURLBuilder_String(t *testing.T) {

	q := url.Values{}
	q.Set("foo", "bar")

	type fields struct {
		base     string
		paths    []string
		query    url.Values
		rawQuery string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "base url only",
			fields: fields{
				base:     "https://www.example.org",
				paths:    nil,
				query:    nil,
				rawQuery: "",
			},
			want: "https://www.example.org",
		},
		{
			name: "base url with paths",
			fields: fields{
				base:     "https://www.example.org",
				paths:    []string{"foo", "bar"},
				query:    nil,
				rawQuery: "",
			},
			want: "https://www.example.org/foo/bar",
		},
		{
			name: "Base url, paths, and query",
			fields: fields{
				base:     "https://www.example.org",
				paths:    []string{"foo", "bar"},
				query:    q,
				rawQuery: "",
			},
			want: "https://www.example.org/foo/bar?foo=bar",
		},
		{
			name: "Base url, paths, and query",
			fields: fields{
				base:     "https://www.example.org",
				paths:    []string{"foo", "bar"},
				query:    q,
				rawQuery: "baz=baz",
			},
			want: "https://www.example.org/foo/bar?baz=baz&foo=bar",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := URLBuilder{
				base:     tt.fields.base,
				paths:    tt.fields.paths,
				query:    tt.fields.query,
				rawQuery: tt.fields.rawQuery,
			}
			if got := b.String(); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}
