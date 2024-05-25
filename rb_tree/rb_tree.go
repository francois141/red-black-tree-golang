package rb_tree

import (
	"errors"
	"golang.org/x/exp/constraints"
)

const (
	RED = iota
	BLACK
)

type RBNode[T constraints.Ordered] struct {
	Data   *T
	Color  int
	Left   *RBNode[T]
	Right  *RBNode[T]
	Parent *RBNode[T]
}

type RBTree[T constraints.Ordered] struct {
	Root    *RBNode[T]
	NILNode *RBNode[T]
}

func (rbTree *RBTree[T]) InitializeNullNode() *RBNode[T] {
	return &RBNode[T]{
		Data:   nil,
		Parent: nil,
		Left:   rbTree.NILNode,
		Right:  rbTree.NILNode,
		Color:  RED,
	}
}

func (rbTree *RBTree[T]) InorderTraversal(node *RBNode[T]) {
	if node == rbTree.NILNode {
		return
	}

	rbTree.InorderTraversal(node.Left)
	//fmt.Printf("%d ", *node.Data)
	rbTree.InorderTraversal(node.Right)
}

func (rbTree *RBTree[T]) LeftRotate(x *RBNode[T]) {
	y := x.Right
	x.Right = y.Left

	if y.Left != nil {
		y.Left.Parent = x
	}
	y.Parent = x.Parent

	if x.Parent == nil {
		rbTree.Root = y
	} else if x == x.Parent.Left {
		x.Parent.Left = y
	} else {
		x.Parent.Right = y
	}

	y.Left = x
	x.Parent = y
}

func (rbTree *RBTree[T]) RightRotate(x *RBNode[T]) {
	y := x.Left
	x.Left = y.Right

	if y.Right != nil {
		y.Right.Parent = x
	}

	y.Parent = x.Parent
	if x.Parent == nil {
		rbTree.Root = y
	} else if x.Parent.Right == x {
		x.Parent.Right = y
	} else {
		x.Parent.Left = y
	}

	y.Right = x
	x.Parent = y
}

func (rbTree *RBTree[T]) Insert(value T) {
	if rbTree.Find(value) {
		return
	}
	newNode := rbTree.InitializeNullNode()
	newNode.Data = &value

	x := rbTree.Root
	var y *RBNode[T] = nil

	for x != rbTree.NILNode {
		y = x
		if *newNode.Data < *x.Data {
			x = x.Left
		} else {
			x = x.Right
		}
	}

	newNode.Parent = y

	if y == nil {
		rbTree.Root = newNode
	} else if *newNode.Data < *y.Data {
		y.Left = newNode
	} else {
		y.Right = newNode
	}

	if newNode.Parent == nil {
		newNode.Color = BLACK
		return
	}
	if newNode.Parent.Parent == nil {
		return
	}

	rbTree.insertFix(newNode)
}

func (rbTree *RBTree[T]) insertFix(node *RBNode[T]) {
	for node.Parent.Color == RED {
		// Right case
		if node.Parent == node.Parent.Parent.Right {
			uncle := node.Parent.Parent.Left
			// Case 1) We recolor only
			if uncle.Color == RED {
				uncle.Color = BLACK
				node.Parent.Color = BLACK
				node.Parent.Parent.Color = RED
				node = node.Parent.Parent
			} else {
				// Case 2)
				if node == node.Parent.Left {
					node = node.Parent
					rbTree.RightRotate(node)
				}
				// Case 2-3)
				node.Parent.Color = BLACK
				node.Parent.Parent.Color = RED
				rbTree.LeftRotate(node.Parent.Parent)
			}
		} else {
			uncle := node.Parent.Parent.Right
			// Case 1)
			if uncle.Color == RED {
				uncle.Color = BLACK
				node.Parent.Color = BLACK
				node.Parent.Parent.Color = RED
				node = node.Parent.Parent
			} else {
				// Case 2)
				if node == node.Parent.Right {
					node = node.Parent
					rbTree.LeftRotate(node)
				}
				// Case 2-3)
				node.Parent.Color = BLACK
				node.Parent.Parent.Color = RED
				rbTree.RightRotate(node.Parent.Parent)
			}
		}
		if node == rbTree.Root {
			break
		}
	}

	rbTree.Root.Color = BLACK
}

func (rbTree *RBTree[T]) Find(key T) bool {
	return rbTree.find(rbTree.Root, key)
}

func (rbTree *RBTree[T]) find(node *RBNode[T], key T) bool {
	if node == nil || node.Data == nil {
		return false
	}
	if *node.Data == key {
		return true
	}

	if key < *node.Data {
		return rbTree.find(node.Left, key)
	} else {
		return rbTree.find(node.Right, key)
	}
}

func (rbTree *RBTree[T]) Size() int {
	return rbTree.size(rbTree.Root)
}

func (rbTree *RBTree[T]) size(node *RBNode[T]) int {
	if node == rbTree.NILNode {
		return 0
	}

	return 1 + rbTree.size(node.Left) + rbTree.size(node.Right)
}

func (rbTree *RBTree[T]) Delete(key T) error {
	return rbTree.delete(rbTree.Root, key)
}

func (rbTree *RBTree[T]) delete(node *RBNode[T], key T) error {
	var z *RBNode[T] = nil

	for node != nil {
		if *node.Data == key {
			z = node
			break
		} else if *node.Data < key {
			node = node.Right
		} else {
			node = node.Left
		}
	}

	if z == nil {
		return errors.New("element not in red black tree")
	}

	var y, x *RBNode[T]
	originalColor := z.Color

	y = z

	if z.Left == rbTree.NILNode {
		x = z.Right
		rbTree.rbTransplant(z, z.Right)
	} else if z.Right == rbTree.NILNode {
		x = z.Left
		rbTree.rbTransplant(z, z.Left)
	} else {
		y = rbTree.minimum(z.Right)
		originalColor = y.Color
		x = y.Right
		if y.Parent == z {
			x.Parent = y
		} else {
			rbTree.rbTransplant(y, y.Right)
			y.Right = z.Right
			y.Right.Parent = y
		}

		rbTree.rbTransplant(z, y)
		y.Left = z.Left
		y.Left.Parent = y
		y.Color = z.Color
	}

	if originalColor == BLACK {
		rbTree.fixDelete(x)
	}

	return nil
}

func (rbTree *RBTree[T]) rbTransplant(u *RBNode[T], v *RBNode[T]) {
	if u.Parent == nil {
		rbTree.Root = v
	} else if u == u.Parent.Left {
		u.Parent.Left = v
	} else {
		u.Parent.Right = v
	}
	v.Parent = u.Parent
}

func (rbTree *RBTree[T]) fixDelete(x *RBNode[T]) {
	var w *RBNode[T]

	for x != rbTree.Root && x.Color == BLACK {
		if x == x.Parent.Left {
			w = x.Parent.Right
			// Case 1
			if w.Color == RED {
				w.Color = BLACK
				x.Parent.Color = RED
				rbTree.LeftRotate(x.Parent)
				w = x.Parent.Right
			}
			// Case 2
			if w.Left.Color == BLACK && w.Right.Color == BLACK {
				w.Color = RED
				x = x.Parent
			} else {
				// Case 3
				if w.Right.Color == BLACK {
					w.Left.Color = BLACK
					w.Color = RED
					rbTree.RightRotate(w)
					w = x.Parent.Right
				}
				// Case 4
				w.Color = x.Parent.Color
				x.Parent.Color = BLACK
				w.Right.Color = BLACK
				rbTree.LeftRotate(x.Parent)
				x = rbTree.Root
			}

		} else if x == x.Parent.Right {
			w = x.Parent.Left
			// Case 1
			if w.Color == RED {
				w.Color = BLACK
				x.Parent.Color = RED
				rbTree.RightRotate(x.Parent)
				w = x.Parent.Left
			}
			// Case 2
			if w.Right.Color == BLACK && w.Left.Color == BLACK {
				w.Color = RED
				x = x.Parent
			} else {
				// Case 3
				if w.Left.Color == BLACK {
					w.Right.Color = BLACK
					w.Color = RED
					rbTree.LeftRotate(w)
					w = x.Parent.Left
				}
				// Case 4
				w.Color = x.Parent.Color
				x.Parent.Color = BLACK
				w.Left.Color = BLACK
				rbTree.RightRotate(x.Parent)
				x = rbTree.Root
			}
		}
	}

	x.Color = BLACK
}

func (rbTree *RBTree[T]) minimum(node *RBNode[T]) *RBNode[T] {
	for node.Left != rbTree.NILNode {
		node = node.Left
	}
	return node
}

func (rbTree *RBTree[T]) maximum(node *RBNode[T]) *RBNode[T] {
	for node.Right != rbTree.NILNode {
		node = node.Right
	}
	return node
}

func NewEmptyRBTree[T constraints.Ordered]() RBTree[T] {
	NILNode := &RBNode[T]{
		Color: BLACK,
		Left:  nil,
		Right: nil,
	}

	return RBTree[T]{
		Root:    NILNode,
		NILNode: NILNode,
	}
}
