package main

import "fmt"

func (t *BinaryTree) SymetricTree(root *TreeNode) {
	var symetric bool
	if root != nil {
		symetric = t.IsSymetric(root.left, root.right)
	} else {
		fmt.Println("Tree is Symetric")
	}

	if symetric {
		fmt.Println("Tree is Symetric")
	} else {
		fmt.Println("Tree is not Symetric")
	}
}
func (t *BinaryTree) IsSymetric(root *TreeNode, root2 *TreeNode) bool {
	if root == nil && root2 == nil {
		return root == root2
	}
	return root.val == root2.val && t.IdenticalTrees(root.left, root2.right) && t.IdenticalTrees(root.right, root2.left)
}
