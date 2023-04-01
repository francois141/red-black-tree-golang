package rb_tree

import "fmt"

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
	Root *RBNode[T]
}

func InitializeNullNode[T comparable](parent *RBNode[T]) *RBNode[T] {
	return &RBNode[T]{
		Data:   0,
		Parent: parent,
		Left:   nil,
		Right:  nil,
		Color:  RED,
	}
}

func (rbTree *RBTree[T]) InorderTraversal(node *RBNode[T]) {
	if node == nil {
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
	newNode := InitializeNullNode[T](nil)
	newNode.Data = value

	x := rbTree.Root
	var y *RBNode[T] = nil

	for x != nil {
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

func NewEmptyRBTree[T comparable]() RBTree[T] {
	return RBTree[T]{
		Root: nil,
	}
}
