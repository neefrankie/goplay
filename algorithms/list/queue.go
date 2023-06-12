package list

type Queue[T interface{}] struct {
	list *DoubleEndedList[T]
}

func NewQueue[T interface{}]() *Queue[T] {
	return &Queue[T]{
		list: NewDoubleEndedList[T](),
	}
}

func (q *Queue[T]) IsEmpty() bool {
	return q.list.IsEmpty()
}

func (q *Queue[T]) Enqueue(item T) {
	q.list.InsertLast(item)
}

func (q *Queue[T]) Dequeue() T {
	return q.list.DeleteFirst()
}
