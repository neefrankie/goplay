package stackque

import "golang.org/x/exp/constraints"

type PriorityQueue[T constraints.Ordered] struct {
	max   int
	arr   []T // Array in sorted, from max at 0 to min at max-1
	count int
}

func NewPriorityQueue[T constraints.Ordered](cap int) *PriorityQueue[T] {
	return &PriorityQueue[T]{
		max:   cap,
		arr:   make([]T, cap),
		count: 0,
	}
}

func (q *PriorityQueue[T]) Insert(item T) {
	if q.count == 0 {
		q.arr[q.count] = item
		q.count++
		return
	}

	// Loop backward.
	i := q.count - 1
	for ; i >= 0; i-- {
		// Shift smaller items upward to make room for item-to-insert.
		if item > q.arr[i] {
			q.arr[i+1] = q.arr[i]
		} else {
			break
		}
	}

	// i is always one slot before the position to insert
	q.arr[i+1] = item
	q.count++
}

func (q *PriorityQueue[T]) Remove() T {
	q.count--
	return q.arr[q.count]
}

func (q *PriorityQueue[T]) PeekMin() T {
	return q.arr[q.count-1]
}

func (q *PriorityQueue[T]) IsEmpty() bool {
	return q.count == 0
}

func (q *PriorityQueue[T]) IsFull() bool {
	return q.count == q.max
}
