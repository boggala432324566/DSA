package main

import "fmt"

func (t *BinaryTree) InorderTraversalRecursion(root *TreeNode) {
	if root == nil {
		return
	}
	t.InorderTraversalRecursion(root.left)
	fmt.Printf("%d ", root.val)
	t.InorderTraversalRecursion(root.right)
}
