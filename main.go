package main

import (
	"fmt"
	"francois141/rbtree/rb_tree"
)

func main() {
	fmt.Println("===================================")
	fmt.Println("= Creation of the red black tree  =")
	fmt.Println("===================================")

	rbTree := rb_tree.NewEmptyRBTree[int]()

	rbTree.Insert(4)
	rbTree.Insert(2)
	rbTree.Insert(5)
	rbTree.Insert(6)

	fmt.Println("===================================")
	fmt.Println("= Traversal of the red black tree =")
	fmt.Println("===================================")

	rbTree.InorderTraversal(rbTree.Root)

	fmt.Println()
}
