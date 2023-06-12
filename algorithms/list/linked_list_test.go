package list

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLinkList(t *testing.T) {
	list := NewLinkedList[float64]()

	list.InsertFirst(2.99)
	list.InsertFirst(4.99)
	list.InsertFirst(6.99)
	list.InsertFirst(8.99)

	list.Traverse(func(item float64) {
		t.Log(item)
	})

	item := list.DeleteFirst()

	assert.Equal(t, 8.99, item)
}

type user struct {
	name string
	id   int
}

func TestLinkList_Find(t *testing.T) {
	list := NewLinkedList[user]()

	list.InsertFirst(user{
		name: "glady",
		id:   5,
	})
	list.InsertFirst(user{
		name: "zieme",
		id:   37,
	})
	list.InsertFirst(user{
		name: "rau",
		id:   63,
	})
	list.InsertFirst(user{
		name: "fisher",
		id:   29,
	})

	got := list.Find(func(item user) bool {
		return item.id == 37
	})

	assert.Equal(t, 37, got.id)

	list.Traverse(func(item user) {
		log.Printf("%v", item)
	})
}

func TestLinkList_Delete(t *testing.T) {
	list := NewLinkedList[user]()

	list.InsertFirst(user{
		name: "glady",
		id:   5,
	})
	list.InsertFirst(user{
		name: "zieme",
		id:   37,
	})
	list.InsertFirst(user{
		name: "rau",
		id:   63,
	})
	list.InsertFirst(user{
		name: "fisher",
		id:   29,
	})

	got := list.Delete(func(item user) bool {
		return item.id == 37
	})

	assert.Equal(t, 37, got.id)

	list.Traverse(func(item user) {
		log.Printf("%v", item)
	})
}

func TestLinkedList_Reverse(t *testing.T) {
	list := NewLinkedList[int]()

	list.InsertFirst(20)
	list.InsertFirst(30)
	list.InsertFirst(40)
	list.InsertFirst(50)

	list.Traverse(func(item int) {
		t.Logf("%d", item)
	})

	list.Reverse()

	t.Log("Reversed: ")
	list.Traverse(func(item int) {
		t.Logf("%d", item)
	})
}

func TestDivider(t *testing.T) {
	t.Logf("2/5 is %d", 2/5)
	t.Logf("5/2 is %d", 5/2)
}

func TestNumberCount(t *testing.T) {
	dividend := 555
	count := 1
	mul := 10

	diff := dividend - mul

	for diff >= 0 {
		t.Logf("diff %d", diff)
		count++
		mul *= 10
		diff = diff - mul
	}

	t.Logf("%d has %d numbers", dividend, count)
}
