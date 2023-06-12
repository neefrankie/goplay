package list

type Iterator[T interface{}] struct {
	current  *Link[T]
	previous *Link[T]
	list     *LinkedList[T]
}

func NewIterator[T interface{}](list *LinkedList[T]) *Iterator[T] {
	return &Iterator[T]{
		current:  list.GetFirst(),
		previous: nil,
		list:     list,
	}
}

func (it *Iterator[T]) Reset() {
	it.current = it.list.GetFirst()
	it.previous = nil
}

func (it *Iterator[T]) AtEnd() bool {
	return it.current.next == nil
}

func (it *Iterator[T]) Next() {
	if it.AtEnd() {
		return
	}

	it.previous = it.current
	it.current = it.current.next
}

func (it *Iterator[T]) GetCurrent() *Link[T] {
	return it.current
}

func (it *Iterator[T]) InsertAfter(item T) {
	link := NewLink(item)
	if it.list.IsEmpty() {
		it.list.SetFirst(link)
		it.current = link
	} else {
		link.next = it.current.next
		it.current.next = link
		it.Next()
	}
}

func (it *Iterator[T]) InsertBefore(item T) {
	link := NewLink(item)
	// Beginning of list or empty list.
	if it.previous == nil {
		link.next = it.list.GetFirst()
		it.list.SetFirst(link)
		it.Reset()
	} else {
		link.next = it.previous.next
		it.previous.next = link
		it.current = link
	}
}

func (it *Iterator[T]) DeleteCurrent() T {
	if it.list.IsEmpty() {
		var t T
		return t
	}

	item := it.current.item
	if it.previous == nil {
		it.list.SetFirst(it.current.next)
		it.Reset()
	} else {
		it.previous.next = it.current.next
		// If current is the last element, delete it and wrap around.
		if it.AtEnd() {
			it.Reset()
		} else {
			it.current = it.current.next
		}
	}

	return item
}
