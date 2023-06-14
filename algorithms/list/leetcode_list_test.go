package list

import (
	"reflect"
	"testing"
)

func TestMyLinkedList(t *testing.T) {
	myLinkedList := Constructor()
	myLinkedList.AddAtHead(1)

	myLinkedList.print()

	myLinkedList.AddAtTail(3)

	myLinkedList.print()

	myLinkedList.AddAtIndex(1, 2)

	myLinkedList.print()

	t.Logf("%d", myLinkedList.Get(1))

	myLinkedList.DeleteAtIndex(1)
	myLinkedList.print()

	t.Logf("%d", myLinkedList.Get(1))

	myLinkedList.DeleteAtIndex(0)
	myLinkedList.print()
}

func Test_removeNthFromEnd(t *testing.T) {
	one := ListNode{Val: 1}
	two := ListNode{Val: 2}
	three := ListNode{Val: 3}
	four := ListNode{Val: 4}
	five := ListNode{Val: 5}

	one.Next = &two
	two.Next = &three
	three.Next = &four
	four.Next = &five

	head := removeNthFromEnd(&one, 2)

	head.traverse(func(v int) {
		t.Logf("%d ->", v)
	})

	one = ListNode{Val: 1}

	head = removeNthFromEnd(&one, 1)
	head.traverse(func(v int) {
		t.Logf("\n%d ->", v)
	})

	two = ListNode{Val: 2}
	one.Next = &two

	head = removeNthFromEnd(head, 1)
	head.traverse(func(v int) {
		t.Logf("\n%d ->", v)
	})
}

func Test_removeElements(t *testing.T) {
	type args struct {
		head *ListNode
		val  int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "[1,2,6,3,4,5,6]",
			args: args{
				head: newListNode(6).add(5).add(4).add(3).add(6).add(2).add(1),
				val:  6,
			},
			want: []int{1, 2, 3, 4, 5},
		},
		{
			name: "[7,7,7,7]",
			args: args{
				head: newListNode(7).add(7).add(7).add(7),
				val:  7,
			},
			want: []int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got := removeElements(tt.args.head, tt.args.val)

			if !reflect.DeepEqual(got.toArray(), tt.want) {
				t.Errorf("removeElements() = %v, want %v", got.toArray(), tt.want)
			}
		})
	}
}

func Test_oddEvenList(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "[1,2,3,4,5,6]",
			args: args{
				head: newListNode(6).add(5).add(4).add(3).add(2).add(1),
			},
			want: []int{1, 3, 5, 2, 4, 6},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := oddEvenList(tt.args.head)

			got.traverse(func(v int) {
				t.Logf("%v", v)
			})

			// if !reflect.DeepEqual(got, tt.want) {
			// 	t.Errorf("oddEvenList() = %v, want %v", got, tt.want)
			// }
		})
	}
}
