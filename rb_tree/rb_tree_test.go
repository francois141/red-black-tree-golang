package rb_tree

import (
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
)

const (
	NB_ITERATIONS int = 1000
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

		assert.Equal(t, i, rbTree.Size())
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

	assert.Equal(t, 7, rbTree.Size())
}

func TestTwoTimesSameValue(t *testing.T) {
	rbTree := NewEmptyRBTree[int]()
	assert.NotNil(t, rbTree)

	rbTree.Insert(55)
	rbTree.Insert(55)

	assert.Equal(t, 1, rbTree.Size())
}

func TestWithRandomValues(t *testing.T) {
	rbTree := NewEmptyRBTree[int]()
	assert.NotNil(t, rbTree)

	mp := make(map[int]struct{}, 0)

	for i := 0; i < NB_ITERATIONS; i++ {
		randomValue := rand.Intn(500)

		rbTree.Insert(randomValue)
		mp[randomValue] = struct{}{}

		assert.Equal(t, len(mp), rbTree.Size())
	}

	for key, _ := range mp {
		assert.True(t, rbTree.Find(key))
	}
}

func TestRBTreeSimpleDelete(t *testing.T) {
	rbTree := NewEmptyRBTree[int]()
	assert.NotNil(t, rbTree)

	value := 4

	rbTree.Insert(value)
	found := rbTree.Find(value)
	assert.True(t, found)

	rbTree.Delete(value)
	found = rbTree.Find(value)
	assert.False(t, found)
}

func TestRBRealTraceDelete(t *testing.T) {
	rbTree := NewEmptyRBTree[int]()
	assert.NotNil(t, rbTree)

	rbTree.Insert(55)
	rbTree.Insert(40)
	rbTree.Insert(65)
	rbTree.Insert(60)
	rbTree.Insert(75)
	rbTree.Insert(57)
	rbTree.Insert(58)

	assert.Equal(t, 7, rbTree.Size())

	rbTree.Delete(55)
	rbTree.Delete(40)
	rbTree.Delete(65)
	rbTree.Delete(60)
	rbTree.Delete(75)
	rbTree.Delete(57)
	rbTree.Delete(58)

	assert.Equal(t, 0, rbTree.Size())
}
