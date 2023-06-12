package list

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewQueue(t *testing.T) {
	q := NewQueue[int]()

	q.Enqueue(20)
	q.Enqueue(30)
	q.Enqueue(40)

	i := q.Dequeue()
	assert.Equal(t, 20, i)

	i = q.Dequeue()
	assert.Equal(t, 30, i)

	i = q.Dequeue()
	assert.Equal(t, 40, i)

	assert.Equal(t, true, q.IsEmpty())
}
