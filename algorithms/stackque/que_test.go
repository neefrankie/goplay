package stackque

import (
	"testing"
)

func TestNewQue(t *testing.T) {
	q := NewQue[int](5)

	t.Logf("Is empty %t", q.IsEmpty())
	t.Logf("%v", q)

	q.Enqueue(1)
	t.Logf("%v", q)

	q.Enqueue(2)
	t.Logf("%v", q)

	q.Enqueue(3)
	t.Logf("%v", q)

	q.Enqueue(4)
	t.Logf("%v", q)

	q.Enqueue(5)
	t.Logf("%v", q)

	t.Logf("Dequeue: ")

	q.Dequeue()
	t.Logf("%v", q)

	q.Dequeue()
	t.Logf("%v", q)

	q.Dequeue()
	t.Logf("%v", q)

	q.Dequeue()
	t.Logf("%v", q)

	q.Dequeue()
	t.Logf("%v", q)
}
