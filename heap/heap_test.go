package heap

import (
	oracle2 "francois141/rbtree/oracle"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSimpleHeap(t *testing.T) {
	heap := NewHeap()

	for i := 9; i >= 0; i-- {
		heap.Push(i)
	}

	for i := 0; i < 10; i++ {
		value, err := heap.Pop()
		assert.NoError(t, err)
		assert.Equal(t, i, value)
	}

	_, err := heap.Pop()
	assert.Error(t, err)
}

func TestVsOracle(t *testing.T) {
	runs := 1000000
	heap := NewHeap()
	oracle := oracle2.NewHeapOracle()

	gen := 1

	for i := 0; i < runs; i++ {
		gen = (gen*7 + 3) % 256 // Test with this generate series - there is no particular reason

		heap.Push(gen)
		oracle.Push(gen)
	}

	for i := 0; i < runs; i++ {
		value, err := heap.Pop()
		assert.NoError(t, err)
		oracleValue, _ := oracle.Pop()
		assert.Equal(t, value, oracleValue)
	}
}
