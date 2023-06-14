package list

import (
	"reflect"
	"testing"
)

func TestListNode_traverse(t *testing.T) {
	type fields struct {
		node *ListNode
	}
	type args struct {
		visitor func(v int)
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "traverse",
			fields: fields{
				node: newListNode(1).add(2).add(3).add(4).add(5),
			},
			args: args{
				visitor: func(v int) {
					t.Logf("%d", v)
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.fields.node.traverse(tt.args.visitor)
		})
	}
}

func Test_traverseListNode(t *testing.T) {
	type args struct {
		head    *ListNode
		visitor func(v int)
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "traverse list node",
			args: args{
				head: newListNode(1).add(2).add(3).add(4).add(5),
				visitor: func(v int) {
					t.Logf("%d", v)
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			traverseListNode(tt.args.head, tt.args.visitor)
		})
	}
}

func TestListNode_toArray(t *testing.T) {
	type fields struct {
		node *ListNode
	}
	tests := []struct {
		name   string
		fields fields
		want   []int
	}{
		{
			name: "collect to array",
			fields: fields{
				node: newListNode(1).add(2).add(3).add(4).add(5),
			},
			want: []int{5, 4, 3, 2, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.fields.node.toArray(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ListNode.toArray() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestListNode_count(t *testing.T) {
	type fields struct {
		node *ListNode
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			name: "count",
			fields: fields{
				node: newListNode(1).add(2).add(3).add(4).add(5),
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			if got := tt.fields.node.count(); got != tt.want {
				t.Errorf("ListNode.count() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_countListNode(t *testing.T) {
	type args struct {
		head *ListNode
		n    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "count nodes",
			args: args{
				head: newListNode(1).add(2).add(3).add(4).add(5),
				n:    0,
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countListNode(tt.args.head, tt.args.n); got != tt.want {
				t.Errorf("countListNode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReverseList(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "reverse list",
			args: args{
				head: newListNode(1).add(2).add(3).add(4).add(5),
			},
			want: []int{1, 2, 3, 4, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ReverseList(tt.args.head)

			gotArr := got.toArray()

			if !reflect.DeepEqual(gotArr, tt.want) {
				t.Errorf("ReverseList() = %v, want %v", got, tt.want)
			}
		})
	}
}
