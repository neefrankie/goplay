package html

import "testing"

func TestTag_String(t1 *testing.T) {

	tests := []struct {
		name   string
		fields Tag
		want   string
	}{
		{
			name: "Tag string",
			fields: Tag{
				token:       Token{"div"},
				selfClosing: false,
				attr:        NewAttr().Set("class", "d-flex"),
				children:    nil,
			},
			want: `<div class="d-flex"></div>`,
		},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &Tag{
				token:       tt.fields.token,
				selfClosing: tt.fields.selfClosing,
				attr:        tt.fields.attr,
				children:    tt.fields.children,
			}
			if got := t.String(); got != tt.want {
				t1.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}
