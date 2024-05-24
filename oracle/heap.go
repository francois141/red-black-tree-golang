package oracle

import "errors"

type HeapOracle struct {
	mp map[int]int
}

func NewHeapOracle() *HeapOracle {
	return &HeapOracle{mp: make(map[int]int)}
}

func (o *HeapOracle) Push(val int) {
	o.mp[val]++
}

func (o *HeapOracle) Pop() (int, error) {
	if len(o.mp) == 0 {
		return 0, errors.New("empty heap")
	}

	best := 100000
	for key, _ := range o.mp {
		best = min(best, key)
	}

	o.mp[best]--
	if o.mp[best] == 0 {
		delete(o.mp, best)
	}

	return best, nil
}

func (o *HeapOracle) Size() int {
	return len(o.mp)
}
