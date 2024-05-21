package b_tree

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBTree_Creation(t *testing.T) {
	assert.NotNil(t, NewBTree())
}

func TestBTree_Insertion(t *testing.T) {
	tree := NewBTree()
	assert.NotNil(t, tree)

	for i := 0; i < 1; i++ {
		tree.Insert(i)
	}
}
