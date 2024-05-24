package oracle

type TreeOracle struct {
	set map[int]struct{}
}

func (t *TreeOracle) Insert(key int) {
	t.set[key] = struct{}{}
}

func (t *TreeOracle) Delete(key int) {
	delete(t.set, key)
}

func (t *TreeOracle) Find(key int) bool {
	_, ok := t.set[key]
	return ok
}

func (t *TreeOracle) Size() int {
	return len(t.set)
}
