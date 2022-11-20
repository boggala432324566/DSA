package main

import (
	"fmt"
)

var slice []int

func (t *BinaryTree) PrintRootToNodePath(root *TreeNode, node int) {
	t.RootToNodePath(root, node)
	fmt.Println("Right View:")
	for _, v := range slice {
		fmt.Println(v)
	}
}

func (t *BinaryTree) RootToNodePath(root *TreeNode, node int) bool {
	if root == nil {
		return false
	}
	slice = append(slice, root.val)
	if node == root.val {
		return true
	}
	if t.RootToNodePath(root.right, node) || t.RootToNodePath(root.left, node) {
		return true
	}
	slice = slice[:len(slice)-1]
	return false
}
