package main

import (
	"fmt"
	"math"
)

func (t *BinaryTree) TreeHeight(root *TreeNode) {
	height := t.Height(root)
	fmt.Println(height)
}
func (t *BinaryTree) Height(root *TreeNode) int {
	if root == nil {
		return 0
	}
	lh := t.Height(root.left)
	rh := t.Height(root.right)
	return 1 + int(math.Max(float64(lh), float64(rh)))
}
