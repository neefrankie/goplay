package stackque

import (
	"testing"
)

func TestNewPriorityQueue(t *testing.T) {
	q := NewPriorityQueue[int](5)

	q.Insert(30)
	q.Insert(50)
	q.Insert(10)
	q.Insert(40)
	q.Insert(20)

	for !q.IsEmpty() {
		item := q.Remove()
		t.Logf("%d ", item)
	}
}
