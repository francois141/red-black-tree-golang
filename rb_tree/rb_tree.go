package rb_tree

import (
	"fmt"
)

const (
	RED = iota
	BLACK
)

type RBNode[T comparable] struct {
	Data   int
	Color  int
	Left   *RBNode[T]
	Right  *RBNode[T]
	Parent *RBNode[T]
}

type RBTree[T comparable] struct {
	Root    *RBNode[T]
	NILNode *RBNode[T]
}

func (rbTree *RBTree[T]) InitializeNullNode() *RBNode[T] {
	return &RBNode[T]{
		Data:   0,
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
	fmt.Printf("%d ", node.Data)
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

func (rbTree *RBTree[T]) Insert(value int) {
	if rbTree.Find(value) {
		return
	}
	newNode := rbTree.InitializeNullNode()
	newNode.Data = value

	x := rbTree.Root
	var y *RBNode[T] = nil

	for x != rbTree.NILNode {
		y = x
		if newNode.Data < x.Data {
			x = x.Left
		} else {
			x = x.Right
		}
	}

	newNode.Parent = y

	if y == nil {
		rbTree.Root = newNode
	} else if newNode.Data < y.Data {
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

func (rbTree *RBTree[T]) Find(key int) bool {
	return rbTree.find(rbTree.Root, key)
}

func (rbTree *RBTree[T]) find(node *RBNode[T], key int) bool {
	if node == nil {
		return false
	}
	if node.Data == key {
		return true
	}

	if key < node.Data {
		return rbTree.find(node.Left, key)
	} else {
		return rbTree.find(node.Right, key)
	}
}

func (rbTree *RBTree[T]) Getsize() int {
	return rbTree.getsize(rbTree.Root)
}

func (rbTree *RBTree[T]) getsize(node *RBNode[T]) int {
	if node == rbTree.NILNode {
		return 0
	}

	return 1 + rbTree.getsize(node.Left) + rbTree.getsize(node.Right)
}

func NewEmptyRBTree[T comparable]() RBTree[T] {
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
