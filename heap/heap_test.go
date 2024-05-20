package heap

import (
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

type oracle struct {
	mp map[int]int
}

func (o *oracle) Push(val int) {
	o.mp[val]++
}

func (o *oracle) Pop() int {
	best := 100000
	for key, _ := range o.mp {
		best = min(best, key)
	}

	o.mp[best]--
	if o.mp[best] == 0 {
		delete(o.mp, best)
	}

	return best
}

func TestVsOracle(t *testing.T) {
	runs := 1000000
	heap := NewHeap()
	oracle := oracle{make(map[int]int)}

	gen := 1

	for i := 0; i < runs; i++ {
		gen = (gen*7 + 3) % 256 // Test with this generate series - there is no particular reason

		heap.Push(gen)
		oracle.Push(gen)
	}

	for i := 0; i < runs; i++ {
		value, err := heap.Pop()
		assert.NoError(t, err)
		assert.Equal(t, value, oracle.Pop())
	}
}
