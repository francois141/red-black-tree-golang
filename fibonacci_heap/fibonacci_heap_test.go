package fibonacci_heap

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFiboHeap(t *testing.T) {
	heap := NewFibHeap()

	for i := 1000; i >= 0; i-- {
		heap.Insert(i)
		value, err := heap.Minimum()
		assert.NoError(t, err)
		assert.Equal(t, i, value)
	}
}

func TestFiboExtractMinimum(t *testing.T) {
	heap := NewFibHeap()

	size := 10000000

	for i := 0; i < size; i++ {
		heap.Insert(i)
	}

	for i := 0; i < size; i++ {
		assert.Equal(t, i, heap.ExtractMinimum().key)
	}
}
