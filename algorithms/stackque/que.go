package stackque

type Que[T interface{}] struct {
	max  int
	head int
	tail int
	arr  []T
}

func NewQue[T interface{}](max int) *Que[T] {
	return &Que[T]{
		max:  max,
		head: max,
		tail: 0,
		arr:  make([]T, max+1),
	}
}

func (q *Que[T]) Enqueue(x T) {
	q.arr[q.tail] = x
	q.tail++
	q.tail = q.tail % q.max
}

func (q *Que[T]) Dequeue() T {
	q.head = q.head % q.max
	x := q.arr[q.head]
	q.head++

	return x
}

func (q *Que[T]) IsEmpty() bool {
	return q.head%q.max == q.tail
}
