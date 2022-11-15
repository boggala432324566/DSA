package main

import (
	"fmt"
	"math"
)

func (t *BinaryTree) CheckBalancedBinaryTree(root *TreeNode) {
	isBalanced := t.Check(root)
	if isBalanced == -1 {
		fmt.Println("Tree is not balanced")
	} else {
		fmt.Println("Tree is balanced")
	}

}
func (t *BinaryTree) Check(root *TreeNode) int {
	if root == nil {
		return 0
	}
	lh := t.Height(root.left)
	rh := t.Height(root.right)
	if lh == -1 || rh == -1 {
		return -1
	}
	if math.Abs(float64(lh-rh)) == -1 {
		return -1

	}
	return 1 + int(math.Max(float64(lh), float64(rh)))
}
