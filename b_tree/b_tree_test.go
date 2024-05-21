package b_tree

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSimpleHeap(t *testing.T) {
	tree := NewBTree()
	assert.NotNil(t, tree)
}
