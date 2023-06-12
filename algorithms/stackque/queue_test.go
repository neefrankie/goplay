package stackque

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestArrQueue_Enqueue(t *testing.T) {
	q := NewQueue[int](5)

	err := q.Enqueue(9)
	t.Logf("%v", q)
	assert.NoError(t, err)

	err = q.Enqueue(4)
	t.Logf("%v", q)
	assert.NoError(t, err)

	err = q.Enqueue(7)
	t.Logf("%v", q)
	assert.NoError(t, err)

	err = q.Enqueue(20)
	t.Logf("%v", q)
	assert.NoError(t, err)

	err = q.Enqueue(3)
	t.Logf("%v", q)
	assert.NoError(t, err)

	err = q.Enqueue(6)
	assert.Error(t, err)
}

func TestArrQueue_Dequeue(t *testing.T) {
	q := NewQueue[int](5)

	err := q.Enqueue(9)
	assert.NoError(t, err)

	err = q.Enqueue(4)
	assert.NoError(t, err)

	err = q.Enqueue(7)
	assert.NoError(t, err)

	err = q.Enqueue(20)
	assert.NoError(t, err)

	err = q.Enqueue(3)
	assert.NoError(t, err)

	e, err := q.Dequeue()
	assert.NoError(t, err)
	assert.Equal(t, e, 9)
	t.Logf("%v", q)

	err = q.Enqueue(6)
	assert.NoError(t, err)
	t.Logf("%v", q)
}

func TestQueue_Traverse(t *testing.T) {
	q := NewQueue[int](5)

	_ = q.Enqueue(3)
	_ = q.Enqueue(9)
	_ = q.Enqueue(7)

	t.Log(q)

	q.Traverse(func(item int) {
		t.Logf("%d", item)
	})

	_ = q.Enqueue(20)
	_ = q.Enqueue(40)

	t.Log(q)

	q.Traverse(func(item int) {
		t.Logf("%d", item)
	})

	_, _ = q.Dequeue()
	_, _ = q.Dequeue()
	_ = q.Enqueue(60)

	t.Log(q)

	q.Traverse(func(item int) {
		t.Log(item)
	})
}
