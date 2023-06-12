package tree

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewTree(t *testing.T) {
	tree := NewBST[int]()

	tree.Insert(50)
	tree.Insert(25)
	tree.Insert(75)
	tree.Insert(12)
	tree.Insert(37)
	tree.Insert(43)
	tree.Insert(30)
	tree.Insert(33)
	tree.InsertR(87)
	tree.InsertR(93)
	tree.InsertR(97)

	tree.Traverse(func(item int) {
		t.Logf("%d", item)
	}, InOrder)

	found, ok := tree.FindR(43)
	if !ok {
		t.Error("not found")
		return
	}

	assert.Equal(t, 43, found)
}
