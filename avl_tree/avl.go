package avl_tree

import "golang.org/x/exp/constraints"

type avlNode[T constraints.Ordered] struct {
	left   *avlNode[T]
	right  *avlNode[T]
	value  int
	height int
}

type avl[T constraints.Ordered] struct {
	root *avlNode[T]
}

func New[T constraints.Ordered]() *avl[T] {
	return &avl[T]{
		root: nil,
	}
}

func (avl *avl[T]) Insert(value int) {
	avl.root = avl.insert(avl.root, value)
}

func (avl *avl[T]) insert(current *avlNode[T], value int) *avlNode[T] {
	// Pure insertion phase
	if current == nil {
		return &avlNode[T]{
			left:   nil,
			right:  nil,
			value:  value,
			height: 1,
		}
	} else if current.value > value {
		current.left = avl.insert(current.left, value)
	} else if current.value < value {
		current.right = avl.insert(current.right, value)
	} else {
		return current
	}

	// Update height here
	current.height = 1 + max(avl.getHeight(current.left), avl.getHeight(current.right))

	// Rebalance phase
	balance := avl.getBalanceFactor(current)

	// Left subtree is bigger
	if balance > 1 {
		if value > current.left.value {
			current.left = avl.leftRotate(current.left)
		}
		return avl.rightRotate(current)
	}

	// Right subtree is bigger
	if balance < -1 {
		if value < current.right.value {
			current.right = avl.rightRotate(current.right)
		}
		return avl.leftRotate(current)
	}

	return current
}

func (avl *avl[T]) getHeight(current *avlNode[T]) int {
	if current == nil {
		return 0
	} else {
		return current.height
	}
}

func (avl *avl[T]) getBalanceFactor(current *avlNode[T]) int {
	return avl.getHeight(current.left) - avl.getHeight(current.right)
}

func (avl *avl[T]) leftRotate(current *avlNode[T]) *avlNode[T] {
	// Perform the rotation
	x := current
	y := current.right
	x.right = y.left
	y.left = x

	// Update heights
	x.height = 1 + max(avl.getHeight(x.left), avl.getHeight(x.right))
	y.height = 1 + max(avl.getHeight(y.left), avl.getHeight(y.right))

	// Return y
	return y
}

func (avl *avl[T]) rightRotate(current *avlNode[T]) *avlNode[T] {
	// Perform the rotation
	x := current
	y := current.left
	x.left = y.right
	y.right = x

	// Update the heights
	x.height = 1 + max(avl.getHeight(x.left), avl.getHeight(x.right))
	y.height = 1 + max(avl.getHeight(y.left), avl.getHeight(y.right))
	// Return y
	return y
}

func (avl *avl[T]) Delete(value int) {
	avl.root = avl.delete(avl.root, value)
}

func (avl *avl[T]) nextValue(current *avlNode[T], value int) int {
	for current.left != nil {
		current = current.left
	}

	return current.value
}

func (avl *avl[T]) delete(current *avlNode[T], value int) *avlNode[T] {
	if current == nil {
		return current
	} else if value < current.value {
		current.left = avl.delete(current.left, value)
	} else if value > current.value {
		current.right = avl.delete(current.right, value)
	} else {
		if current.left == nil && current.right == nil {
			// Case 1 - leaf
			return nil
		} else if current.left == nil {
			// Case 2 - one child
			current = current.right
		} else if current.right == nil {
			// Case 2 - one child
			current = current.left
		} else {
			// Case 3 - two childrens
			succ := avl.nextValue(current.right, value)
			current.value = succ
			current.right = avl.delete(current.right, succ)
		}
	}

	// Update height
	current.height = 1 + max(avl.getHeight(current.left), avl.getHeight(current.right))

	// Rebalance phase
	balance := avl.getBalanceFactor(current)
	if balance > 1 {
		if avl.getBalanceFactor(current.left) < 0 {
			current.left = avl.leftRotate(current.left)
		}
		return avl.rightRotate(current)
	}

	// Right subtree is bigger
	if balance < -1 {
		if avl.getBalanceFactor(current.right) > 0 {
			current.right = avl.rightRotate(current.right)
		}
		return avl.leftRotate(current)
	}

	return current
}

func (avl *avl[T]) Find(value int) bool {
	return avl.find(avl.root, value)
}

func (avl *avl[T]) find(current *avlNode[T], value int) bool {
	if current == nil {
		return false
	}

	if value < current.value {
		return avl.find(current.left, value)
	} else if value > current.value {
		return avl.find(current.right, value)
	} else {
		return true
	}
}

func (avl *avl[T]) Size() int {
	return avl.size(avl.root)
}

func (avl *avl[T]) size(current *avlNode[T]) int {
	if current == nil {
		return 0
	}

	return 1 + avl.size(current.left) + avl.size(current.right)
}
