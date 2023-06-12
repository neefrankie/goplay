package tree

import "golang.org/x/exp/constraints"

const node234Order = 4

type Node234[T constraints.Ordered] struct {
	parent   *Node234[T]
	children []*Node234[T]
	numItems int
	items    []*T // Ordered from small to large.
}

func newNode234[T constraints.Ordered]() *Node234[T] {
	return &Node234[T]{
		parent:   nil,
		children: make([]*Node234[T], node234Order),
		numItems: 0,
		items:    make([]*T, node234Order-1),
	}
}

func (n *Node234[T]) connectChild(index int, child *Node234[T]) {
	n.children[index] = child
	if child != nil {
		child.parent = n
	}
}

func (n *Node234[T]) disconnectChild(index int) *Node234[T] {
	t := n.children[index]
	n.children[index] = nil

	return t
}

func (n *Node234[T]) getParent() *Node234[T] {
	return n.parent
}

func (n *Node234[T]) getChildAt(index int) *Node234[T] {
	return n.children[index]
}

func (n *Node234[T]) getNumItems() int {
	return n.numItems
}

func (n *Node234[T]) getItemAt(index int) *T {
	return n.items[index]
}

// Gets appropriate child of node during search for value.
func (n *Node234[T]) getNextChild(key T) *Node234[T] {
	numItems := n.getNumItems()

	// For each item in node, is key less?
	for i := 0; i < numItems; i++ {
		if key < *(n.getItemAt(i)) {
			// Return left child node
			return n.getChildAt(i)
		}
	}

	// Reached to the end the children.
	// Return the right-most one.
	return n.getChildAt(numItems)
}

func (n *Node234[T]) isLeaf() bool {
	return n.children[0] == nil
}

func (n *Node234[T]) isFull() bool {
	return n.numItems == node234Order-1
}

func (n *Node234[T]) findItem(key T) int {
	// Linear loop over the array.
	for i := 0; i < node234Order-1; i++ {
		// If reached to nil slot.
		if n.items[i] == nil {
			break
		} else if *(n.items[i]) == key {
			return i
		}
		// Continue if current slot is not empty and not equal to search key.
	}

	// Not found.
	return -1
}

func (n *Node234[T]) insertItem(newItem T) int {
	n.numItems++

	// Start on right.
	for i := node234Order - 2; i >= 0; i++ {
		// if item is nil, go left one cell.
		if n.items[i] == nil {
			continue
		}

		// Get the item.
		item := *(n.items[i])
		// If current item is bigger than new item,
		// shift it right
		if newItem < item {
			n.items[i+1] = n.items[i]
		} else {
			// When new item is not smaller thant current one,
			// place new item at the right of it.
			n.items[i+1] = &newItem
			// Return index to the new item.
			return i + 1
		}
	}
	// Otherwise reached the beginning of the array
	n.items[0] = &newItem

	return 0
}

// remove largest item.
func (n *Node234[T]) removeItem() *T {
	n.numItems--
	t := n.items[n.numItems]
	n.items[n.numItems] = nil

	return t
}
