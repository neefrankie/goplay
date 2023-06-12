package stackque

import "errors"

// Queue is a first-in, first-out queue implementation based on array.
// The field size keeps track of current valid items in the queue, and tests
// whether the queue is full or empty.
// If you do not use this field, you have to rely on hea and tail to
// figure out whether the queue is empty of full.
// In such case, make the array one cell larger than the maximum number
// of items that will be placed in it.
type Queue[T interface{}] struct {
	arr  []T
	head int
	tail int
	cap  int
	size int // Number of valid items.
}

func NewQueue[T interface{}](cap int) *Queue[T] {
	// cap = cap + 1 to implement a size-free version.
	return &Queue[T]{
		arr:  make([]T, cap),
		head: 0,
		tail: -1,
		cap:  cap,
		size: 0,
	}
}

// Enqueue puts item at end to queue.
func (q *Queue[T]) Enqueue(x T) error {
	if q.IsFull() {
		return errors.New("queue is full")
	}

	// Wrap around
	if q.tail == q.cap-1 {
		q.tail = -1
	}

	// Increment tail and insert one more item
	q.tail++
	q.arr[q.tail] = x
	q.size++

	return nil
}

func (q *Queue[T]) Dequeue() (T, error) {
	if q.IsEmpty() {
		var ret T
		return ret, errors.New("queue is empty")
	}

	t := q.arr[q.head]
	q.head++
	if q.head == q.cap {
		q.head = 0
	}
	q.size--

	return t, nil
}

func (q *Queue[T]) PeekHead() T {
	return q.arr[q.head]
}

func (q *Queue[T]) PeekTail() T {
	return q.arr[q.tail]
}

func (q *Queue[T]) IsEmpty() bool {
	// tail + 1 == head for size-free version
	return q.size == 0
}

func (q *Queue[T]) IsFull() bool {
	// rear + 2 == head || head + cap - 2 == rear
	return q.size == q.cap
}

func (q *Queue[T]) Size() int {
	// if rear >= head return rear - head else return (cap - front) + (rear + 1)
	return q.size
}

// Traverse visit each item from head to tail.
func (q *Queue[T]) Traverse(visitor func(item T)) {
	if q.size == 0 {
		return
	}
	if q.tail >= q.head {
		for i := q.head; i <= q.tail; i++ {
			visitor(q.arr[i])
		}
		return
	}

	for i := q.head; i < q.cap; i++ {
		visitor(q.arr[i])
	}

	for i := 0; i <= q.tail; i++ {
		visitor(q.arr[i])
	}
}
