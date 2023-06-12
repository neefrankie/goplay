package list

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewStack(t *testing.T) {
	s := NewStack[int]()

	s.Push(10)
	s.Push(20)
	s.Push(30)

	i := s.Pop()
	assert.Equal(t, 30, i)

	i = s.Pop()
	assert.Equal(t, 20, i)

	i = s.Pop()
	assert.Equal(t, 10, i)

	assert.Equal(t, true, s.IsEmpty())
}
