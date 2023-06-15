package list

import "fmt"

type MyLinkedList struct {
	head *ListNode
	size int
}

func Constructor() MyLinkedList {
	return MyLinkedList{
		head: nil,
		size: 0,
	}
}

func (list *MyLinkedList) Get(index int) int {
	if index < 0 || index > list.size {
		return -1
	}

	cur := list.head
	for i := 0; i < list.size; i++ {
		cur = cur.Next
	}

	return cur.Val
}

func (list *MyLinkedList) AddAtHead(val int) {
	list.AddAtIndex(0, val)
}

func (list *MyLinkedList) AddAtTail(val int) {
	list.AddAtIndex(list.size, val)
}

// Add a node of value val before the indexth node in the linked list.
// If index equals the length of the linked list,
// the node will be appended to the end of the linked list.
// If index is greater than the length, the node will not be inserted.
func (list *MyLinkedList) AddAtIndex(index int, val int) {
	if index > list.size {
		return
	}

	cur := list.head
	newNode := newListNode(val)

	if index <= 0 {
		newNode.Next = cur
		list.head = newNode
	} else {
		for i := 0; i < index-1; i++ {
			cur = cur.Next
		}
		newNode.Next = cur.Next
		cur.Next = newNode
	}

	list.size = list.size + 1
}

func (list *MyLinkedList) DeleteAtIndex(index int) {
	if index < 0 || index >= list.size {
		return
	}

	cur := list.head

	if index == 0 {
		list.head = list.head.Next
	} else {
		for i := 0; i < index-1; i++ {
			cur = cur.Next
		}
		cur.Next = cur.Next.Next
	}

	list.size = list.size - 1
}

func (list *MyLinkedList) print() {
	cur := list.head

	for cur != nil {
		fmt.Printf("%d -> ", cur.Val)
		cur = cur.Next
	}

	fmt.Println()
}

func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}

	fast := head
	slow := head
	for fast != nil && fast.Next != nil {

		fast = fast.Next.Next
		slow = slow.Next

		if fast == slow {
			return true
		}
	}

	return false
}

// getIntersectionNode returns the node at which the two lists intersect.
// LeetCode linked list 214.
// See https://leetcode.com/explore/learn/card/linked-list/214/two-pointer-technique/1215/discuss/49785/Java-solution-without-knowing-the-difference-in-len!
// See https://leetcode.com/explore/learn/card/linked-list/214/two-pointer-technique/1215/discuss/2116221/Visual-Explanation-or-One-Pass-JAVA
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}

	curA := headA
	curB := headB

	for curA != curB {
		if curA == nil {
			curA = headB
		} else {
			curA = curA.Next
		}

		if curB == nil {
			curB = headA
		} else {
			curB = curB.Next
		}
	}

	return curA
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil {
		return head
	}

	fast := head
	slow := head

	for i := 0; i < n; i++ {
		if fast != nil {
			fast = fast.Next
		}
	}

	// when n is the same as the length of the list.
	// just return head.next without needing to stitch together the two sides of the target node.
	if fast == nil {
		return head.Next
	}

	for fast.Next != nil {
		fast = fast.Next
		slow = slow.Next
	}

	if slow == head && n == 1 {
		return nil
	}

	slow.Next = slow.Next.Next

	return head
}

func removeElements(head *ListNode, val int) *ListNode {
	if head == nil {
		return nil
	}

	dummy := &ListNode{}
	dummy.Next = head
	prev := dummy
	cur := head

	for cur != nil {
		if cur.Val == val {
			prev.Next = cur.Next
		} else {
			prev = cur
		}

		cur = cur.Next
	}

	return dummy.Next
}

func removeElementsRec(head *ListNode, val int) *ListNode {
	if head == nil {
		return nil
	}

	head.Next = removeElementsRec(head.Next, val)

	if head.Val == val {
		return head.Next
	}

	return head
}

func oddEvenList(head *ListNode) *ListNode {
	var oddHead = &ListNode{}
	var oddTail = oddHead
	var evenHead = &ListNode{}
	var evenTail = evenHead

	isOdd := true
	cur := head
	for cur != nil {

		if isOdd {
			oddTail.Next = cur
			oddTail = oddTail.Next
		} else {
			evenTail.Next = cur
			evenTail = evenTail.Next
		}
		isOdd = !isOdd
		cur = cur.Next
	}

	oddTail.Next = evenHead.Next
	evenTail.Next = nil

	return oddHead.Next
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	head := &ListNode{}

	cur := head

	cur1 := list1
	cur2 := list2

	for cur1 != nil && cur2 != nil {
		if cur1.Val < cur2.Val {
			cur.Next = cur1
			cur1 = cur1.Next
		} else {
			cur.Next = cur2
			cur2 = cur2.Next
		}
		cur = cur.Next
	}

	for cur1 != nil {
		cur.Next = cur1
		cur1 = cur1.Next
		cur = cur.Next
	}

	for cur2 != nil {
		cur.Next = cur2
		cur2 = cur2.Next
		cur = cur.Next
	}

	cur.Next = nil
	return head.Next
}

// You are given two non-empty linked lists representing two non-negative integers.
// The digits are stored in reverse order,
// and each of their nodes contains a single digit.
// Add the two numbers and return the sum as a linked list.
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	head := &ListNode{}
	cur := head

	carry := 0
	for l1 != nil || l2 != nil || carry != 0 {
		var sum int
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}

		sum += carry

		carry = sum / 10

		node := &ListNode{Val: sum % 10}
		cur.Next = node
		cur = cur.Next
	}

	return head.Next
}

type Node struct {
	Val    int
	Prev   *Node
	Next   *Node
	Child  *Node
	Random *Node
}

func newNode(v int) *Node {
	return &Node{Val: v}
}

func (node *Node) addNext(n *Node) *Node {
	node.Next = n
	n.Prev = node

	return node
}

func (node *Node) addPrev(n *Node) *Node {
	node.Prev = n
	n.Next = node
	return node
}

func (node *Node) addChild(n *Node) *Node {
	node.Child = n
	return node
}

func (node *Node) traverse(visitor func(n *Node)) {
	cur := node
	for cur != nil {
		visitor(cur)
		if cur.Child != nil {
			cur.Child.traverse(visitor)
		}

		cur = cur.Next
	}
}

func flatten(root *Node) *Node {
	h, _ := flattenRec(root)
	return h
}

func flattenRec(root *Node) (*Node, *Node) {
	tail := root
	cur := root
	for cur != nil {
		if cur.Child == nil {
			tail = cur
			cur = cur.Next
			continue
		}

		h, t := flattenRec(cur.Child)

		t.Next = cur.Next
		if cur.Next != nil {
			cur.Next.Prev = t
		}

		cur.Next = h
		h.Prev = cur

		cur.Child = nil
	}

	return root, tail
}

// Iterative approach.
// See https://leetcode.com/explore/learn/card/linked-list/213/conclusion/1225/discuss/150321/Easy-Understanding-Java-beat-95.7-with-Explanation
func flattenIter(root *Node) *Node {
	if root == nil {
		return nil
	}

	outer := root
	for outer != nil {
		if outer.Child == nil {
			continue
		}

		inner := outer.Child
		for inner.Next != nil {
			inner = inner.Next
		}

		inner.Next = outer.Next
		if outer.Next != nil {
			outer.Next.Prev = inner
		}

		outer.Next = outer.Child
		outer.Child.Prev = outer

		outer.Child = nil
	}

	return root
}

func copyRandomList(head *Node) *Node {
	cur := head

	// First round: make copy of each node,
	// and link them together side-by-side in a single list.
	for cur != nil {
		origNext := cur.Next

		copy := Node{Val: cur.Val}
		cur.Next = &copy
		copy.Next = origNext

		cur = origNext
	}

	// Second round: assign random pointers for the copy nodes.
	cur = head
	for cur != nil {
		if cur.Random != nil {
			cur.Next.Random = cur.Random.Next
		}

		cur = cur.Next.Next
	}

	// Third round: restore the original list, and extract the copy list.
	cur = head
	dummy := &Node{}
	copyCur := dummy

	for cur != nil {
		origNext := cur.Next.Next
		copy := cur.Next

		copyCur.Next = copy
		copyCur = copyCur.Next

		cur.Next = origNext
		cur = cur.Next
	}

	return dummy.Next
}

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	cur := head
	size := 1
	for cur.Next != nil {
		size++
		cur = cur.Next
	}

	cur.Next = head

	k = k % size

	gap := size - k - 1

	cur = head
	for i := gap; i > 0; i-- {
		cur = cur.Next
	}

	newHead := cur.Next
	cur.Next = nil

	return newHead
}
