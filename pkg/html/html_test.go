package html

import (
	"testing"
)

func TestNewHTML(t *testing.T) {
	type args struct {
		lang string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "HTML",
			args: args{
				lang: "en",
			},
			want: `<!DOCTYPE html><html lang="en"></html>`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewHTML(tt.args.lang)

			if got.String() != tt.want {
				t.Errorf("NewHTML() = %v, want %v", got, tt.want)
			}
		})
	}
}
