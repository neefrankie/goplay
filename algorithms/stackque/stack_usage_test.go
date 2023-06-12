package stackque

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReverseWord(t *testing.T) {
	type args struct {
		word string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Reverse hello",
			args: args{
				word: "Hello",
			},
			want: "olleH",
		},
		{
			name: "Reverse part",
			args: args{
				word: "part",
			},
			want: "trap",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, ReverseWord(tt.args.word), "ReverseWord(%v)", tt.args.word)
		})
	}
}

func TestCheckBracket(t *testing.T) {
	str := "a{b(c]d}e)"

	err := CheckBracket(str, func(err *BracketError) {
		if err != nil {
			t.Error(err)
		}
	})

	if err != nil {
		t.Error(err)
	}
}
