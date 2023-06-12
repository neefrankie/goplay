package list

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
