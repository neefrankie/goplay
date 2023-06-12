package tree

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBinaryNode_CountDuplicate(t *testing.T) {
	root := NewBinaryNode[byte]('E', nil, nil)

	root.InsertR('A')
	root.InsertR('S')
	root.InsertR('Y')
	root.InsertR('Q')
	root.InsertR('U')
	root.InsertR('E')
	root.InsertR('S')
	root.InsertR('T')
	root.InsertR('I')
	root.InsertR('O')
	root.InsertR('N')

	root.InOrderTraverse(func(item byte) {
		t.Logf("%s", string(item))
	})

	assert.Equal(t, 2, root.CountDuplicate('E'))
	assert.Equal(t, 2, root.CountDuplicate('S'))
}
