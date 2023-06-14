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
