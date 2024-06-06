package fibonacci_heap

type FibHeap struct {
	n        int
	min      *FibNode
	rootList *FibNode
}

// TODO: Make generic
type FibNode struct {
	key    int
	degree int
	mark   bool

	parent *FibNode
	child  *FibNode
	left   *FibNode
	right  *FibNode
}

func (fibHeap *FibHeap) Insert(key int) {
	// Create a node
	node := &FibNode{key: key}
	node.left = node
	node.right = node

	fibHeap.insertRootList(node)

	// Check if this is the new minimum
	if fibHeap.min == nil || key < fibHeap.min.key {
		fibHeap.min = node
	}

	// Update size of the tree
	fibHeap.n++
}

func (fibHeap *FibHeap) insertRootList(node *FibNode) {
	if fibHeap.rootList == nil {
		fibHeap.rootList = node
	} else {
		node.left = fibHeap.rootList.left
		node.right = fibHeap.rootList
		fibHeap.rootList.left.right = node
		fibHeap.rootList.left = node
	}
}

func (fibHeap *FibHeap) Minimum() int {
	return fibHeap.min.key
}

func (fibHeap *FibHeap) Union(other *FibHeap) {
	// Get minimums
	if other.min.key < fibHeap.min.key {
		fibHeap.min = other.min
	}

	// Combine root lists
	oldLeft := fibHeap.rootList.left

	fibHeap.rootList.left.right = other.rootList
	fibHeap.rootList.left = other.rootList.left

	other.rootList.left.right = fibHeap.rootList
	other.rootList.left = oldLeft

	// Update count
	fibHeap.n += other.n
}
