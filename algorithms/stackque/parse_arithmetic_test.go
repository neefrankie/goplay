package stackque

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPostfixBuilder_Convert(t *testing.T) {

	tests := []struct {
		name string
		arg  string
		want string
	}{
		{
			name: "2+3",
			arg:  "2+3",
			want: "2 3 +",
		},
		{
			name: "2+3",
			arg:  "2 + 3  * 4",
			want: "2 3 4 * +",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := NewPostfixBuilder(tt.arg)
			assert.Equalf(t, tt.want, b.Convert(), "Convert()")
		})
	}
}

func Test_evaluatePostfix(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "3*(4+5)-6/(1+2)",
			args: args{
				input: "3 4 5 + * 6 1 2 + / -",
			},
			want: 25,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, evaluatePostfix(tt.args.input), "evaluatePostfix(%v)", tt.args.input)
		})
	}
}

func Test_readNum(t *testing.T) {
	str := "45 32 +"

	end := readNum(str, 0)

	num := str[0:end]

	assert.Equal(t, "45", num)
}

func Test_readNumStack(t *testing.T) {
	stack := NewStack[int](10)

	str := "45 32 +"

	_ = stack.Push(0)
	for i := 0; str[i] >= '0' && str[i] <= '9'; i++ {
		prev, _ := stack.Pop()
		curr := 10*prev + int(str[i]-'0')
		_ = stack.Push(curr)
	}

	num, _ := stack.Pop()

	assert.Equal(t, 45, num)
}
