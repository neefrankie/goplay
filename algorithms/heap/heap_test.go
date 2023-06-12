package heap

import (
	"testing"
)

func TestNew(t *testing.T) {
	pq := New()

	pq.Insert(70)
	pq.Insert(40)
	pq.Insert(50)
	pq.Insert(20)
	pq.Insert(60)
	pq.Insert(100)
	pq.Insert(80)
	pq.Insert(30)
	pq.Insert(10)
	pq.Insert(90)

	t.Logf("%v", pq.heapArray)
}
