package main

import "fmt"

func (t *BinaryTree) PreorderTraversalRecursion(root *TreeNode) {
	if root == nil {
		return
	}
	fmt.Printf("%d ", root.val)
	t.PreorderTraversalRecursion(root.left)
	t.PreorderTraversalRecursion(root.right)
}
