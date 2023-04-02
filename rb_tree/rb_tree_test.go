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

func TestRBTreeSize(t *testing.T) {
	rbTree := NewEmptyRBTree[int]()
	assert.NotNil(t, rbTree)

	for i := 1; i <= 1000; i++ {
		rbTree.Insert(i)

		found := rbTree.Find(i)
		assert.True(t, found)

		assert.Equal(t, i, rbTree.Getsize())
	}
}

func TestRBRealTrace(t *testing.T) {
	rbTree := NewEmptyRBTree[int]()
	assert.NotNil(t, rbTree)

	rbTree.Insert(55)
	rbTree.Insert(40)
	rbTree.Insert(65)
	rbTree.Insert(60)
	rbTree.Insert(75)
	rbTree.Insert(57)
	rbTree.Insert(58)

	assert.Equal(t, 7, rbTree.Getsize())
}
