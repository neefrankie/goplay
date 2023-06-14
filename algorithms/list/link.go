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

func Reverse[T interface{}](link *Link[T]) *Link[T] {

	var current = link
	var previous *Link[T]

	for current != nil {
		temp := current.next
		current.next = previous
		previous = current
		current = temp
	}

	return previous
}
