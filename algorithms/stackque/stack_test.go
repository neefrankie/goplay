package stackque

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewStack(t *testing.T) {
	stack := NewStack[int](5)
	isEmpty := stack.IsEmpty()
	assert.Equal(t, true, isEmpty)

	_ = stack.Push(2)
	_ = stack.Push(4)
	_ = stack.Push(9)
	_ = stack.Push(38)
	_ = stack.Push(26)

	top, _ := stack.Pop()
	assert.Equal(t, 26, top)

	_ = stack.Push(42)
	isFull := stack.IsFull()
	assert.Equal(t, true, isFull)

	stack.Traverse(func(item int) {
		t.Logf("%d", item)
	})
}

func TestStack_Push(t *testing.T) {

	stack := NewStack[int](5)

	type args struct {
		x int
	}
	tests := []struct {
		name    string
		args    args
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "Push1",
			args: args{
				x: 9,
			},
			wantErr: assert.NoError,
		},
		{
			name: "Push2",
			args: args{
				x: 36,
			},
			wantErr: assert.NoError,
		},
		{
			name: "Push3",
			args: args{
				x: 18,
			},
			wantErr: assert.NoError,
		},
		{
			name: "Push4",
			args: args{
				x: 23,
			},
			wantErr: assert.NoError,
		},
		{
			name: "Push5",
			args: args{
				x: 84,
			},
			wantErr: assert.NoError,
		},
		{
			name: "Push to full",
			args: args{
				x: 55,
			},
			wantErr: assert.Error,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			tt.wantErr(t, stack.Push(tt.args.x), fmt.Sprintf("Push(%v)", tt.args.x))
		})
	}
}
