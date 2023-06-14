package list

import "fmt"

type MyLinkedList struct {
	head *ListNode
}

func Constructor() MyLinkedList {
	return MyLinkedList{}
}

func (this *MyLinkedList) Get(index int) int {
	cur := this.head
	i := 0
	for cur != nil && i < index {
		cur = cur.Next
		i++
	}

	if cur != nil {
		return cur.Val
	}

	return -1
}

func (this *MyLinkedList) AddAtHead(val int) {

	node := ListNode{Val: val, Next: nil}
	node.Next = this.head

	this.head = &node
}

func (this *MyLinkedList) AddAtTail(val int) {

	if this.head == nil {
		node := ListNode{Val: val}
		this.head = &node
		return
	}

	cur := this.head
	prev := this.head

	for cur != nil {
		prev = cur
		cur = cur.Next
	}

	node := ListNode{Val: val}

	prev.Next = &node
}

// Add a node of value val before the indexth node in the linked list.
// If index equals the length of the linked list,
// the node will be appended to the end of the linked list.
// If index is greater than the length, the node will not be inserted.
func (this *MyLinkedList) AddAtIndex(index int, val int) {
	if index == 0 {
		node := ListNode{Val: val}
		node.Next = this.head
		this.head = &node
		return
	}

	cur := this.head
	prev := this.head

	i := 0
	for cur != nil && i < index {
		prev = cur
		cur = cur.Next
		i++
	}

	if i < index {
		return
	}

	node := ListNode{Val: val}
	node.Next = cur
	prev.Next = &node
}

func (this *MyLinkedList) DeleteAtIndex(index int) {
	if index == 0 {
		if this.head == nil {
			return
		}
		this.head = this.head.Next
		return
	}

	prev := this.head
	cur := this.head
	i := 0

	for cur != nil && i < index {
		prev = cur
		cur = cur.Next
		i++
	}

	if cur != nil {
		prev.Next = cur.Next
	}
}

func (this *MyLinkedList) print() {
	cur := this.head

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
