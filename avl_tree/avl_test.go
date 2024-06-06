package avl_tree

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNew(t *testing.T) {
	assert.NotNil(t, New())
}

func TestInsert(t *testing.T) {
	tree := New()

	size := 2000
	for i := 0; i < size; i++ {
		tree.Insert(i)
		assert.True(t, tree.Find(i))
		assert.False(t, tree.Find(i+1))
	}
}

func TestDelete(t *testing.T) {
	tree := New()

	size := 20000
	for i := 0; i < size; i++ {
		tree.Insert(i)
	}

	for i := 0; i < size; i++ {
		tree.Delete(i)
		assert.False(t, tree.Find(i))
	}
}

func TestSize(t *testing.T) {
	tree := New()

	size := 20000
	for i := 0; i < size; i++ {
		tree.Insert(i)
		tree.Insert(i)
	}

	for i := 0; i < size; i++ {
		tree.Delete(i)
		assert.False(t, tree.Find(i))
	}

	for i := 1; i <= size; i++ {
		tree.Insert(i)
		assert.Equal(t, i, tree.Size())
	}
}
