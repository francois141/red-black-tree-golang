package b_tree

import "golang.org/x/exp/constraints"

const (
	t int = 3
)

type BTreeNode[T constraints.Ordered] struct {
	isLeaf    bool
	childrens []*BTreeNode[T]
	keys      []int
}

func NewBTreeNode[T constraints.Ordered]() *BTreeNode[T] {
	return &BTreeNode[T]{
		isLeaf:    true,
		childrens: make([]*BTreeNode[T], 0, 2*t-1),
		keys:      make([]int, 0, 2*t),
	}
}

type BTree[T constraints.Ordered] struct {
	root *BTreeNode[T]
}

func NewBTree[T constraints.Ordered]() *BTree[T] {
	return &BTree[T]{
		root: NewBTreeNode[T](),
	}
}

func (tree *BTree[T]) _splitNode(current *BTreeNode[T], position int, child *BTreeNode[T]) {
	// Allocate new node
	otherChild := NewBTreeNode[T]()
	otherChild.isLeaf = child.isLeaf

	// Copy from child / keys to new node
	for i := t; i < 2*t-1; i++ {
		otherChild.keys = append(otherChild.keys, child.keys[i])
	}

	if !otherChild.isLeaf {
		for i := t; i < 2*t; i++ {
			otherChild.childrens = append(otherChild.childrens, child.childrens[i])
		}
	}

	// Remove values from the original child
	medianValue := child.keys[t-1]
	child.keys = child.keys[:t-1]
	child.childrens = child.childrens[:t]

	// Alter state of x
	current.keys = append(current.keys, 0)
	for i := len(current.keys) - 2; i >= position; i-- {
		current.keys[i+1] = current.keys[i]
	}

	current.childrens = append(current.childrens, NewBTreeNode[T]())
	for i := len(current.childrens) - 2; i >= position+1; i-- {
		current.childrens[i+1] = current.childrens[i]
	}

	// Finally insert to current
	current.keys[position] = medianValue
	current.childrens[position+1] = otherChild
}

func (tree *BTree[T]) Insert(key int) {
	// Check if we need to split the root first
	if len(tree.root.keys) == 2*t-1 {
		// Get old root
		oldRoot := tree.root
		newRoot := NewBTreeNode[T]()
		newRoot.isLeaf = false
		newRoot.childrens = append(newRoot.childrens, oldRoot)

		// Split the root using old root as child
		tree._splitNode(newRoot, 0, oldRoot)

		// Set the new root
		tree.root = newRoot
	}

	// If root is empty
	if len(tree.root.keys) == 0 {
		tree.root.keys = append(tree.root.keys, key)
		return
	}
	// Invariant: root isn't full here
	tree._insert(tree.root, key)
}

func (tree *BTree[T]) _insert(current *BTreeNode[T], key int) {
	if current.isLeaf {
		idx := len(current.keys) - 1
		current.keys = append(current.keys, 0)
		for current.keys[idx] > key {
			current.keys[idx+1] = current.keys[idx]
			idx--
		}
		current.keys[idx+1] = key
	} else {
		idx := 0
		for idx < len(current.keys) && current.keys[idx] < key {
			idx++
		}
		child := current.childrens[idx]
		if len(child.keys) == 2*t-1 {
			tree._splitNode(current, idx, child)
			if key > current.keys[idx] {
				child = current.childrens[idx+1]
			}
		}
		tree._insert(child, key)
	}
}

func (tree *BTree[T]) Find(key int) bool {
	return tree.find(tree.root, key)
}

func (tree *BTree[T]) find(current *BTreeNode[T], key int) bool {
	for _, value := range current.keys {
		if value == key {
			return true
		}
	}

	if !current.isLeaf {
		idx := 0
		for idx < len(current.keys) && current.keys[idx] < key {
			idx++
		}

		return tree.find(current.childrens[idx], key)
	}

	return false
}

func (tree *BTree[T]) Delete(key int) {
	tree.delete(tree.root, key)
}

func (tree *BTree[T]) delete(current *BTreeNode[T], key int) {
	idx := tree.findKey(current, key)

	if idx < len(current.keys) && current.keys[idx] == key {
		if current.isLeaf {
			tree.removeFromLeaf(current, idx)
		} else {
			tree.removeFromNonLeaf(current, idx)
		}

		return
	}

	// Case 3 - node isn't present - make sure child has t-1 values at least
	if current.isLeaf {
		return
	}

	if len(current.childrens[idx].keys) < t {
		tree.fill(current, idx)
	}

	// If we are in the last ==> reduce len by -1 as we squashed it with merge
	tree.delete(current.childrens[min(idx, len(current.childrens)-1)], key)
}

func (tree *BTree[T]) fill(current *BTreeNode[T], idx int) {
	if idx > 0 && len(current.childrens[idx-1].keys) >= t {
		tree.borrowFromPred(current, idx)
		return
	}

	if idx < len(current.childrens)-1 && len(current.childrens[idx+1].keys) >= t {
		tree.borrowFromSucc(current, idx)
		return
	}
}

func (tree *BTree[T]) borrowFromPred(current *BTreeNode[T], idx int) {
	child := current.childrens[idx]
	prev := current.childrens[idx-1]

	// Borrow
	child.keys = append([]int{current.keys[idx-1]}, child.keys...)
	if len(prev.childrens) > 0 {
		child.childrens = append([]*BTreeNode[T]{prev.childrens[len(prev.childrens)-1]}, child.childrens...)
	}

	// Change key of parent
	current.keys[idx-1] = prev.keys[len(prev.keys)-1]

	// Remove values from prev
	if len(prev.keys) > 0 {
		prev.keys = prev.keys[:len(prev.keys)-1]
	}

	if len(prev.childrens) > 0 {
		prev.childrens = prev.childrens[:len(prev.childrens)-1]
	}
}

func (tree *BTree[T]) borrowFromSucc(current *BTreeNode[T], idx int) {
	child := current.childrens[idx]
	next := current.childrens[idx+1]

	// Borrow
	child.keys = append(child.keys, current.keys[idx])
	if len(next.childrens) > 0 {
		child.childrens = append(child.childrens, next.childrens[0])
	}

	// Change the key of the parent
	current.keys[idx] = next.keys[0]

	// Remove values from next
	if len(next.keys) >= 1 {
		next.keys = next.keys[1:]
	}

	if len(next.childrens) >= 1 {
		next.childrens = next.childrens[1:]
	}
}

func (tree *BTree[T]) findKey(current *BTreeNode[T], key int) int {
	idx := 0
	for idx < len(current.keys) && current.keys[idx] < key {
		idx++
	}

	return idx
}

func (tree *BTree[T]) removeFromLeaf(current *BTreeNode[T], idx int) {
	if !current.isLeaf {
		panic("Should be leaf")
	}

	for i := idx + 1; i < len(current.keys); i++ {
		current.keys[i-1] = current.keys[i]
	}
	current.keys = current.keys[:len(current.keys)-1]
}

func (tree *BTree[T]) getPredecessor(current *BTreeNode[T]) int {
	for !current.isLeaf {
		current = current.childrens[len(current.childrens)-1]
	}

	return current.keys[len(current.keys)-1]
}

func (tree *BTree[T]) getSuccessor(current *BTreeNode[T]) int {
	for !current.isLeaf {
		current = current.childrens[0]
	}

	return current.keys[0]
}

func (tree *BTree[T]) removeFromNonLeaf(current *BTreeNode[T], idx int) {
	if current.isLeaf {
		panic("Should be non leaf")
	}

	// Case 2a - left has at least t nodes - swap values and recursive delete
	if len(current.childrens[idx].childrens) > t {
		value := tree.getPredecessor(current.childrens[idx+1])
		current.keys[idx] = value
		tree.delete(current.childrens[idx], value)
		return
	}

	// Case 2b - right has at least t nodes - swap values and recursive delete
	if len(current.childrens[idx+1].childrens) > t {
		value := tree.getSuccessor(current.childrens[idx+1])
		current.keys[idx] = value
		tree.delete(current.childrens[idx+1], value)
		return
	}

	// Case 2c
	key := current.keys[idx]
	tree.merge(current, idx)
	tree.delete(current, key)
}

func (tree *BTree[T]) merge(current *BTreeNode[T], idx int) {
	// Merge both childrens together
	c1 := current.childrens[idx]
	c2 := current.childrens[idx+1]

	// Merge keys
	c1.keys = append(c1.keys, current.keys[idx])
	c1.keys = append(c1.keys, c2.keys...)

	// Merge childrens
	// TODO: Is it correct?
	c1.childrens = append(c1.childrens, c2.childrens...)

	// Alter now the state of current
	current.keys = append(current.keys[:idx], current.keys[idx+1:]...)
	current.childrens = append(current.childrens[:idx], current.childrens[idx+1:]...)
}

func (tree *BTree[T]) Size() int {
	return tree.size(tree.root)
}

func (tree *BTree[T]) size(current *BTreeNode[T]) int {
	if current == nil {
		return 0
	}

	output := len(current.keys)

	for _, child := range current.childrens {
		output += tree.size(child)
	}

	return output
}
