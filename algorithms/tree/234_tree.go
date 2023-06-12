package tree

import "golang.org/x/exp/constraints"

type T234[T constraints.Ordered] struct {
	root *Node234[T]
}

func NewTree234[T constraints.Ordered]() *T234[T] {
	return &T234[T]{
		root: newNode234[T](),
	}
}

func (t *T234[T]) Find(key T) int {
	curNode := t.root

	for {
		// Found it in current node.
		if index := curNode.findItem(key); index != -1 {
			return index
		} else if curNode.isLeaf() {
			// Not found in current node, and current node has
			// no children, we cannot go any further.
			return -1
		} else {
			// Not found in current node; however, there are
			// child node, and we can search deeper.
			curNode = curNode.getNextChild(key)
		}
	}
}

func (t *T234[T]) Insert(item T) {
	curNode := t.root

	for {
		if curNode.isFull() {
			t.split(curNode)
			// Back up to parent to restart search.
			curNode = curNode.getParent()
			curNode = curNode.getNextChild(item)
		} else if curNode.isLeaf() {
			break
		} else {
			curNode = curNode.getNextChild(item)
		}
	}

	curNode.insertItem(item)
}

// Node split:
// Assume the data items are named A, B, and C.
// - A new, empty node is created. It's a sibling of the node being split,
// and is placed to its right.
// - Date item C is moved into the new node.
// - Date item B is moved into the parent of the node being split.
// - Date item A remains where it is.
// - The rightmost two children are disconnected from the node
// being split and connected to the new node.
// Additionally, if the node to split is root, a new node is created to
// become the parent of the node being split.
func (t *T234[T]) split(node *Node234[T]) {
	// Assumes node is full.
	// Remove items from this node.
	itemC := node.removeItem()
	itemB := node.removeItem()

	// Remove children from this node.
	child2 := node.disconnectChild(2)
	child3 := node.disconnectChild(3)
	var parent *Node234[T]

	newRight := newNode234[T]()

	// If this is the root
	if node == t.root {
		t.root = newNode234[T]()
		parent = t.root
		// The old root becomes the first child of the new root.
		t.root.connectChild(0, node)
		// The new root has no items upto this point.
	} else {
		parent = node.getParent()
	}

	// Insert item B to parent.
	bIndex := parent.insertItem(*itemB)
	n := parent.getNumItems()

	// Move parent's connections one child to the right.
	// Loop backward from end to inserted position.
	for i := n - 1; i > bIndex; i-- {
		temp := parent.disconnectChild(i)
		parent.connectChild(i+1, temp)
	}

	// Connect newRight to parent.
	parent.connectChild(bIndex+1, newRight)

	// Item c to newRight.
	newRight.insertItem(*itemC)
	// Connect to 0 and 1 on newRight.
	newRight.connectChild(0, child2)
	newRight.connectChild(1, child3)
}
