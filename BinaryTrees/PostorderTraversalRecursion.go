package main

import "fmt"

func (t *BinaryTree) PostorderTraversalRecursion(root *TreeNode) {
	if root == nil {
		return
	}
	t.PostorderTraversalRecursion(root.left)
	t.PostorderTraversalRecursion(root.right)
	fmt.Printf("%d ", root.val)
}