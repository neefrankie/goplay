package list

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
