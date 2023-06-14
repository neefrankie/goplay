package list

type ListNode struct {
	Val  int
	Next *ListNode
}

func newListNode(v int) *ListNode {
	return &ListNode{Val: v}
}

func (l *ListNode) add(v int) *ListNode {
	head := ListNode{Val: v}
	head.Next = l

	return &head
}

func (l *ListNode) traverse(visitor func(v int)) {
	cur := l

	for cur != nil {
		visitor(cur.Val)
		cur = cur.Next
	}
}

func traverseListNode(head *ListNode, visitor func(v int)) {
	if head == nil {
		return
	}

	visitor(head.Val)
	traverseListNode(head.Next, visitor)
}

func (l *ListNode) toArray() []int {
	out := make([]int, 0)
	l.traverse(func(v int) {
		out = append(out, v)
	})

	return out
}

func (l *ListNode) count() int {
	c := 0

	l.traverse(func(v int) {
		c++
	})

	return c
}

func countListNode(head *ListNode, n int) int {
	if head == nil {
		return n
	}

	n++
	return countListNode(head.Next, n)
}

func ReverseList(head *ListNode) *ListNode {
	var prev *ListNode
	cur := head

	for cur != nil {
		next := cur.Next
		cur.Next = prev

		prev = cur
		cur = next
	}

	return prev
}
