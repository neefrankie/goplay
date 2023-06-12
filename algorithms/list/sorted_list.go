package list

// SortedList arranges items in order.
// Deletion is limited to the smallest or largest item in the list,
// which is at the start of the list.
type SortedList[T interface{}] struct {
	head *Link[T]
	less func(current T, newItem T) bool // When true, current is placed before newItem
}

func NewSortedList[T interface{}](less func(a T, b T) bool) *SortedList[T] {
	return &SortedList[T]{
		head: nil,
		less: less,
	}
}

func (list *SortedList[T]) IsEmpty() bool {
	return list.head == nil
}

func (list *SortedList[T]) Insert(newItem T) {
	link := NewLink(newItem)

	var previous *Link[T] = nil // newItem should be inserted after previous.
	current := list.head

	for current != nil && list.less(current.item, newItem) {
		previous = current
		current = current.next
	}

	// Should be placed at the beginning, including empty list.
	if previous == nil {
		list.head = link
	} else {
		previous.next = link
	}

	link.next = current
}

func (list *SortedList[T]) Remove() T {
	if list.IsEmpty() {
		var t T
		return t
	}

	t := list.head
	list.head = list.head.next

	return t.item
}

func (list *SortedList[T]) Traverse(visitor func(item T)) {
	current := list.head
	if current == nil {
		return
	}

	for current != nil {
		visitor(current.item)
		current = current.next
	}
}
