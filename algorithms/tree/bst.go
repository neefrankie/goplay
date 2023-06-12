package tree

import "golang.org/x/exp/constraints"

type BinarySearchTree[T constraints.Ordered] struct {
	root *BinaryNode[T]
}

func NewBST[T constraints.Ordered]() *BinarySearchTree[T] {
	return &BinarySearchTree[T]{
		root: nil,
	}
}

func (t *BinarySearchTree[T]) FindR(key T) (T, bool) {
	return bstSearchR(t.root, key)
}

// Find is a non-recursive version of search.
func (t *BinarySearchTree[T]) Find(key T) T {
	// Hod the node currently examining.
	current := t.root

	for current != nil {
		if key < current.item { // Go left
			current = current.left
		} else if key > current.item { // Go right
			current = current.right
		} else {
			break
		}
	}

	if current == nil {
		var t T
		return t
	}

	return current.item
}

func (t *BinarySearchTree[T]) InsertR(item T) {
	t.root = bstInsertR(t.root, item)
}

// Insert is a non-recursive version to insert a node.
// To insert a node, we must first find the place to insert it.
func (t *BinarySearchTree[T]) Insert(item T) {
	node := NewBinaryNode(item, nil, nil)
	if t.root == nil {
		t.root = node
	} else {
		current := t.root
		// The last non-null node encountered.
		// This is necessary because `current` is set to `nil`
		// in the process of discovering that its parent did not have
		// an appropriate child.
		// If we didn't save `parent`, we would lose track of where we were.
		var parent *BinaryNode[T]

		for {
			parent = current
			if item < current.item {
				current = current.left
				if current == nil {
					parent.left = node
					return
				}
			} else {
				current = current.right
				if current == nil {
					parent.right = node
					return
				}
			}
		}
	}
}

func (t *BinarySearchTree[T]) Delete(key T) bool {
	current := t.root
	parent := t.root
	isLeftChild := true

	for current.item != key {
		parent = current
		if key < current.item {
			isLeftChild = true
			current = current.left
		} else {
			isLeftChild = false
			current = current.right
		}

		if current == nil {
			return false
		}
	}

	// If no children, simply delete it
	if current.left == nil && current.right == nil {
		if current == t.root {
			t.root = nil
		} else if isLeftChild {
			parent.left = nil
		} else {
			parent.right = nil
		}

		return true
	}

	// If no right child.
	if current.right == nil {
		if current == t.root {
			t.root = current.left
		} else if isLeftChild {
			parent.left = current.left
		} else {
			parent.right = current.left
		}
		return true
	}

	// If no left child.
	if current.left == nil {
		if current == t.root {
			t.root = current.right
		} else if isLeftChild {
			parent.left = current.right
		} else {
			parent.right = current.right
		}

		return true
	}

	// Two children, replace with in-order successor
	successor := t.getSuccessor(current)
	if current == t.root {
		t.root = successor
	} else if isLeftChild {
		parent.left = successor
	} else {
		parent.right = successor
	}
	// Connect successor to current's left child.
	// Successor cannot have a left child.
	successor.left = current.left
	return true
}

// getSuccessor finds the next-highest value after delNode.
// Goes to right child, then right child's left descendants.
func (t *BinarySearchTree[T]) getSuccessor(delNode *BinaryNode[T]) *BinaryNode[T] {
	parent := delNode
	successor := delNode
	current := delNode.right

	for current != nil {
		parent = successor
		successor = current
		current = current.left
	}

	// Current is nil now, which means successor only has an optional right child
	if successor != delNode.right {
		parent.left = successor.right
		successor.right = delNode.right
	}

	return successor
}

func (t *BinarySearchTree[T]) Traverse(visitor func(item T), order TraverseType) {
	switch order {
	case PreOrder:
		preOrderTraversal(t.root, visitor)

	case InOrder:
		inOrderTraversal(t.root, visitor)

	case PostOrder:
		inOrderTraversal(t.root, visitor)
	}
}
