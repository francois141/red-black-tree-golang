package fibonacci_heap

import (
	oracle2 "francois141/rbtree/oracle"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFiboHeap(t *testing.T) {
	heap := NewFibHeap()

	for i := 1000; i >= 0; i-- {
		heap.Push(i)
		value, err := heap.Pop()
		assert.NoError(t, err)
		assert.Equal(t, i, value)
	}
}

func TestFiboExtractMinimum(t *testing.T) {
	heap := NewFibHeap()

	size := 1000000

	for i := 0; i < size; i++ {
		heap.Push(i)
	}

	for i := 0; i < size; i++ {
		key, err := heap.Pop()
		assert.NoError(t, err)
		assert.Equal(t, i, key)
	}

	_, err := heap.Pop()
	assert.Error(t, err)
}

func TestMultipleSameValue(t *testing.T) {
	heap := NewFibHeap()

	size := 1000000

	for i := 0; i < size; i++ {
		heap.Push(0)
	}

	for i := 0; i < size; i++ {
		key, err := heap.Pop()
		assert.NoError(t, err)
		assert.Equal(t, 0, key)
	}

	_, err := heap.Pop()
	assert.Error(t, err)
}

func TestVsOracleFibo(t *testing.T) {
	runs := 1000000
	heap := NewFibHeap()
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
