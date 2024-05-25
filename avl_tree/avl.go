package avl_tree

// TODO: Make this data structure generic
type avlNode struct {
	left   *avlNode
	right  *avlNode
	value  int
	height int
}

type avl struct {
	root *avlNode
}

func New() *avl {
	return &avl{
		root: &avlNode{},
	}
}

func (avl *avl) Insert(value int) {
	avl.root = avl.insert(avl.root, value)
}

func (avl *avl) insert(current *avlNode, value int) *avlNode {
	// Pure insertion phase
	if current == nil {
		return &avlNode{
			left:   nil,
			right:  nil,
			value:  value,
			height: 1,
		}
	} else if current.value > value {
		current.left = avl.insert(current.left, value)
	} else if current.value < value {
		current.right = avl.insert(current.right, value)
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

func (avl *avl) getHeight(current *avlNode) int {
	if current == nil {
		return 0
	} else {
		return current.height
	}
}

func (avl *avl) getBalanceFactor(current *avlNode) int {
	return avl.getHeight(current.left) - avl.getHeight(current.right)
}

func (avl *avl) leftRotate(current *avlNode) *avlNode {
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

func (avl *avl) rightRotate(current *avlNode) *avlNode {
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

func (avl *avl) Delete(value int) {
	avl.root = avl.delete(avl.root, value)
}

func (avl *avl) nextValue(current *avlNode, value int) int {
	for current.left != nil {
		current = current.left
	}

	return current.value
}

func (avl *avl) delete(current *avlNode, value int) *avlNode {
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

func (avl *avl) Find(value int) bool {
	return avl.find(avl.root, value)
}

func (avl *avl) find(current *avlNode, value int) bool {
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
