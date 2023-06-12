package list

type Link[T interface{}] struct {
	item T
	next *Link[T]
}

func NewLink[T interface{}](item T) *Link[T] {
	return &Link[T]{
		item: item,
	}
}

func (l *Link[T]) FindMiddle() *Link[T] {
	middle := l
	end := l

	for end != nil && end.next != nil {
		middle = middle.next
		end = end.next.next
	}

	return middle
}

type DoublyLink[T interface{}] struct {
	item     T
	next     *DoublyLink[T]
	previous *DoublyLink[T]
}

func NewDoublyLink[T interface{}](item T) *DoublyLink[T] {
	return &DoublyLink[T]{
		item: item,
	}
}
