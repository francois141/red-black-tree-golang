package fibonacci_heap

import (
	"errors"
	"math"
)

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

func NewFibHeap() *FibHeap {
	return &FibHeap{}
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

func (fibHeap *FibHeap) Minimum() (int, error) {
	if fibHeap.rootList == nil {
		return 0, errors.New("empty fibonacci heap")
	}
	return fibHeap.min.key, nil
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

func (fibHeap *FibHeap) ExtractMinimum() *FibNode {
	if fibHeap.min == nil {
		return nil
	}

	minimumNode := fibHeap.min

	// Remove children
	for fibHeap.min.child != nil {
		child := fibHeap.min.child
		fibHeap.min.child = fibHeap.removeFromList(fibHeap.min.child)
		fibHeap.insertRootList(child)
	}

	fibHeap.removeRootList(minimumNode)

	// Remove node from list
	if minimumNode.left == minimumNode {
		fibHeap.min = nil
		fibHeap.rootList = nil
	} else {
		fibHeap.min = minimumNode.right
		fibHeap.consolidate()
	}

	fibHeap.n--

	return minimumNode
}

func (fibHeap *FibHeap) consolidate() {
	buffer := make([]*FibNode, int(math.Log(float64(fibHeap.n))*2))
	nodes := fibHeap.getRoots()

	for _, node := range nodes {
		degree := node.degree
		for buffer[degree] != nil {
			other := buffer[degree]
			if node.key > other.key {
				node, other = other, node
			}

			// Link other node to the child
			fibHeap.heapLink(other, node)

			buffer[degree] = nil
			degree++
		}

		buffer[degree] = node
	}

	for _, node := range buffer {
		if node != nil && node.key < fibHeap.min.key {
			fibHeap.min = node
		}
	}
}

func (fibHeap *FibHeap) getRoots() []*FibNode {
	current := fibHeap.rootList
	first := current
	output := make([]*FibNode, 0)
	output = append(output, current)

	for current.right != first {
		output = append(output, current.right)
		current = current.right
	}

	return output
}

func (fibHeap *FibHeap) heapLink(other *FibNode, node *FibNode) {
	// Step 1: Remove other from node
	fibHeap.removeRootList(other)
	other.left = other
	other.right = other

	fibHeap.insertIntoChildList(node, other)

	// Step 2 add node to other
	node.degree++
	other.parent = node
	other.mark = false
}

func (fibHeap *FibHeap) removeRootList(node *FibNode) {
	if fibHeap.rootList == node {
		fibHeap.rootList = node.right
	}

	fibHeap.removeFromList(node)
}

func (fibHeap *FibHeap) removeFromList(node *FibNode) *FibNode {
	if node.left == node {
		return nil
	}

	if node != nil {
		node.left.right = node.right
		node.right.left = node.left

		return node.right
	}

	return nil
}

func (fibHeap *FibHeap) insertIntoChildList(parent *FibNode, child *FibNode) {
	if parent.child == nil {
		parent.child = child
	} else {
		child.right = parent.child.right
		child.left = parent.child
		parent.child.right.left = child
		parent.child.right = child
	}
}
