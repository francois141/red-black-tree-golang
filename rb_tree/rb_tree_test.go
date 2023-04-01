package rb_tree

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateEmptyRBTree(t *testing.T) {
	rbTree := NewEmptyRBTree[int]()
	assert.NotNil(t, rbTree)
}

func TestRBTreeInsertSingleValue(t *testing.T) {
	rbTree := NewEmptyRBTree[int]()
	assert.NotNil(t, rbTree)

	value := 4

	rbTree.Insert(value)
	found := rbTree.Find(value)
	assert.True(t, found)
}
