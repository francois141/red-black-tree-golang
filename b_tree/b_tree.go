package b_tree

const (
	t int = 3
)

// TODO: I could make the data structure generic
type BTreeNode struct {
	isLeaf    bool
	childrens []*BTreeNode
	keys      []int
}

func NewBTreeNode() *BTreeNode {
	return &BTreeNode{
		isLeaf:    true,
		childrens: make([]*BTreeNode, 0, 2*t-1),
		keys:      make([]int, 0, 2*t),
	}
}

type BTree struct {
	root *BTreeNode
}

func NewBTree() *BTree {
	return &BTree{
		root: NewBTreeNode(),
	}
}

func (tree *BTree) _splitNode(current *BTreeNode, position int, child *BTreeNode) {
	// Allocate new node
	otherChild := NewBTreeNode()
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

	current.childrens = append(current.childrens, NewBTreeNode())
	for i := len(current.childrens) - 2; i >= position+1; i-- {
		current.childrens[i+1] = current.childrens[i]
	}

	// Finally insert to current
	current.keys[position] = medianValue
	current.childrens[position+1] = otherChild
}

func (tree *BTree) Insert(key int) {
	// Check if we need to split the root first
	if len(tree.root.keys) == 2*t-1 {
		// Get old root
		oldRoot := tree.root
		newRoot := NewBTreeNode()
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

func (tree *BTree) _insert(current *BTreeNode, key int) {
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

func (tree *BTree) Find(key int) bool {
	return tree._find(tree.root, key)
}

func (tree *BTree) _find(current *BTreeNode, key int) bool {
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

		return tree._find(current.childrens[idx], key)
	}

	return false
}
