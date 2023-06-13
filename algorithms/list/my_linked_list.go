package list

import "fmt"

type Node struct {
	val  int
	next *Node
}

type MyLinkedList struct {
	head *Node
}

func Constructor() MyLinkedList {
	return MyLinkedList{}
}

func (this *MyLinkedList) Get(index int) int {
	cur := this.head
	i := 0
	for cur != nil && i < index {
		cur = cur.next
		i++
	}

	if cur != nil {
		return cur.val
	}

	return -1
}

func (this *MyLinkedList) AddAtHead(val int) {

	node := Node{val: val, next: nil}
	node.next = this.head

	this.head = &node
}

func (this *MyLinkedList) AddAtTail(val int) {

	if this.head == nil {
		node := Node{val: val}
		this.head = &node
		return
	}

	cur := this.head
	prev := this.head

	for cur != nil {
		prev = cur
		cur = cur.next
	}

	node := Node{val: val}

	prev.next = &node
}

// Add a node of value val before the indexth node in the linked list.
// If index equals the length of the linked list,
// the node will be appended to the end of the linked list.
// If index is greater than the length, the node will not be inserted.
func (this *MyLinkedList) AddAtIndex(index int, val int) {
	if index == 0 {
		node := Node{val: val}
		node.next = this.head
		this.head = &node
		return
	}

	cur := this.head
	prev := this.head

	i := 0
	for cur != nil && i < index {
		prev = cur
		cur = cur.next
		i++
	}

	if i < index {
		return
	}

	node := Node{val: val}
	node.next = cur
	prev.next = &node
}

func (this *MyLinkedList) DeleteAtIndex(index int) {
	if index == 0 {
		if this.head == nil {
			return
		}
		this.head = this.head.next
		return
	}

	prev := this.head
	cur := this.head
	i := 0

	for cur != nil && i < index {
		prev = cur
		cur = cur.next
		i++
	}

	if cur != nil {
		prev.next = cur.next
	}
}

func (this *MyLinkedList) print() {
	cur := this.head

	for cur != nil {
		fmt.Printf("%d -> ", cur.val)
		cur = cur.next
	}

	fmt.Println()
}
