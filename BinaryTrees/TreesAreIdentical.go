package main

import (
	"fmt"
)

func (t *BinaryTree) TreesAreIdentical(root *TreeNode, root2 *TreeNode) {
	identical := t.IdenticalTrees(root, root2)
	if identical {
		fmt.Println("Tree are identical")
	} else {
		fmt.Println("Tree are not identical")
	}
}
func (t *BinaryTree) IdenticalTrees(root *TreeNode, root2 *TreeNode) bool {
	if root == nil && root2 == nil {
		return root == root2
	}
	return root.val == root2.val && t.IdenticalTrees(root.left, root2.left) && t.IdenticalTrees(root.right, root2.right)
}
