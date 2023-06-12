package tree

import (
	"golang.org/x/exp/constraints"
)

type TraverseType int

const (
	_ TraverseType = iota
	PreOrder
	InOrder
	PostOrder
)

type BinaryNode[T constraints.Ordered] struct {
	item  T
	left  *BinaryNode[T]
	right *BinaryNode[T]
}

func NewBinaryNode[T constraints.Ordered](item T, left, right *BinaryNode[T]) *BinaryNode[T] {
	return &BinaryNode[T]{
		item:  item,
		left:  left,
		right: right,
	}
}

// bstSearchR searches a binary tree recursively.
func bstSearchR[T constraints.Ordered](node *BinaryNode[T], key T) (T, bool) {
	if node == nil {
		var t T
		return t, false
	}

	if key == node.item {
		return node.item, true
	}

	if key < node.item {
		// rotateRight to bring the target node to top
		return bstSearchR(node.left, key)
	} else {
		// rotateLeft to bring to top.
		return bstSearchR(node.right, key)
	}
}

func bstInsertR[T constraints.Ordered](node *BinaryNode[T], item T) *BinaryNode[T] {
	if node == nil {
		return NewBinaryNode(item, nil, nil)
	}

	if item < node.item {
		node.left = bstInsertR(node.left, item)
	} else {
		node.right = bstInsertR(node.right, item)
	}

	return node
}

// insertTop inserts performs root insertion.
func bstInsertTop[T constraints.Ordered](node *BinaryNode[T], item T) *BinaryNode[T] {
	if node == nil {
		return NewBinaryNode(item, nil, nil)
	}

	if item < node.item {
		node.left = bstInsertTop(node.left, item)
		rotateRight(node)
	} else {
		node.right = bstInsertTop(node.right, item)
		rotateLeft(node)
	}

	return node
}

// Splay insertion brings newly inserted nodes to the root with
// two rotations.
func bstSplay[T constraints.Ordered](node *BinaryNode[T], item T) *BinaryNode[T] {
	if node == nil {
		return NewBinaryNode(item, nil, nil)
	}

	// Should go left
	if item < node.item {
		// Current node has no left child, set it to new node's
		// right child
		if node.left == nil {
			return NewBinaryNode(item, nil, node)
		}

		// Go left-left: rotate right at the root twice.
		if item < node.left.item {
			node.left.left = bstSplay(node.left.left, item)
			node = rotateRight(node)
		} else {
			// Go left-right: rotate left at the left child,
			// then right at the root
			node.left.right = bstSplay(node.left.right, item)
			node.left = rotateLeft(node.left)
		}

		return rotateRight(node)
	} else {
		// Show go right
		if node.right == nil {
			return NewBinaryNode(item, node, nil)
		}

		// Go right-right: rotate left at the root twice
		if node.right.item < item {
			node.right.right = bstSplay(node.right.right, item)
			node = rotateLeft(node)
		} else {
			// Go right-left: rotate right at the right child,
			// then left at the root.
			node.right.left = bstSplay(node.right.left, item)
			node.right = rotateRight(node.right)
		}

		return rotateLeft(node)
	}
}

func bstDeleteR[T constraints.Ordered](node *BinaryNode[T], key T) *BinaryNode[T] {
	if node == nil {
		return nil
	}

	if key < node.item {
		bstDeleteR(node.left, key)
	} else {
		bstDeleteR(node.right, key)
	}

	if key == node.item {
		// Replace current node with the next largest in order one.
		node = bubbleNextInOrder(node.right)
	}

	return node
}

func bubbleNextInOrder[T constraints.Ordered](root *BinaryNode[T]) *BinaryNode[T] {
	if root.left == nil {
		return root
	}

	root.left = bubbleNextInOrder(root.left)
	root = rotateRight(root)

	return root
}

// Combine two subtrees into one tree.
// To combine two BSTs with all keys in the second known tobe larger than all keys in the first,
// bring the smallest element in that tree to the root.
// At this point, the left subtree of the root must be empty.
func joinLR[T constraints.Ordered](left *BinaryNode[T], right *BinaryNode[T]) *BinaryNode[T] {
	if right == nil {
		return left
	}

	right = bubbleNextInOrder(right)
	right.left = left

	return right
}

func rotateRight[T constraints.Ordered](oldRoot *BinaryNode[T]) *BinaryNode[T] {
	newRoot := oldRoot.left
	// Copy the right node of the left child to be the
	// left node of the old root.
	oldRoot.left = newRoot.right
	newRoot.right = oldRoot

	return newRoot
}

func rotateLeft[T constraints.Ordered](oldRoot *BinaryNode[T]) *BinaryNode[T] {
	newRoot := oldRoot.right
	oldRoot.right = newRoot.left
	newRoot.left = oldRoot

	return newRoot
}

func (n *BinaryNode[T]) DeleteR(key T) {
	bstDeleteR(n, key)
}

func (n *BinaryNode[T]) InsertR(item T) *BinaryNode[T] {
	return bstInsertR(n, item)
}

func (n *BinaryNode[T]) CountDuplicate(item T) int {
	current := n
	var count int

	for current != nil {
		if item < current.item {
			current = current.left
		} else {
			if item == current.item {
				count++
			}
			current = current.right
		}
	}

	return count
}

func (n *BinaryNode[T]) PreOrderTraverse(visit func(item T)) {
	preOrderTraversal(n, visit)
}

func (n *BinaryNode[T]) InOrderTraverse(visit func(item T)) {
	inOrderTraversal(n, visit)
}

func (n *BinaryNode[T]) PostOrderTraverse(visit func(item T)) {
	postOrderTraversal(n, visit)
}

// preOrderTraversal visits a node, then visit the left and right subtrees.
func preOrderTraversal[T constraints.Ordered](localRoot *BinaryNode[T], visitor func(item T)) {
	if localRoot != nil {
		visitor(localRoot.item)
		preOrderTraversal(localRoot.left, visitor)
		preOrderTraversal(localRoot.right, visitor)
	}
}

// inOrderTraversal visits the left subtree, then visit the node, then visit the right subtree.
func inOrderTraversal[T constraints.Ordered](localRoot *BinaryNode[T], visitor func(item T)) {
	if localRoot != nil {
		inOrderTraversal(localRoot.left, visitor)
		visitor(localRoot.item)
		inOrderTraversal(localRoot.right, visitor)
	}
}

// postOrderTraversal visits the left and right subtree, then visit the node.
func postOrderTraversal[T constraints.Ordered](localRoot *BinaryNode[T], visitor func(item T)) {
	if localRoot != nil {
		postOrderTraversal(localRoot.left, visitor)
		postOrderTraversal(localRoot.right, visitor)
		visitor(localRoot.item)
	}
}
