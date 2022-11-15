package main

import "fmt"

func (t *BinaryTree) BoundaryTraversal(root *TreeNode) {
	var boundaryTra []int
	boundaryTra = append(boundaryTra, root.val)
	LeftNodes(root, &boundaryTra)
	LeafNodes(root, &boundaryTra)
	RightNodes(root, &boundaryTra)
	for _, v := range boundaryTra {
		fmt.Println(v)
	}
}

func LeftNodes(root *TreeNode, slice *[]int) {
	curr := root.left
	for curr != nil {
		if !isLeaf(root) {
			*slice = append(*slice, root.val)
		}
		if root.left != nil {
			curr = root.left
		} else {
			curr = root.right
		}
	}

}
func LeafNodes(root *TreeNode, slice *[]int) {
	if root == nil {
		return
	}
	if isLeaf(root) {
		*slice = append(*slice, root.val)
	}
	LeafNodes(root.left, slice)
	LeafNodes(root.right, slice)
}

func RightNodes(root *TreeNode, slice *[]int) {
	var rightSlice []int
	curr := root.right
	for curr != nil {
		if !isLeaf(root) {
			rightSlice = append(rightSlice, root.val)
		}
		if root.right != nil {
			curr = root.right
		} else {
			curr = root.left
		}
	}

	for i := len(rightSlice) - 1; i >= 0; i-- {
		*slice = append(*slice, rightSlice[i])
	}

}

func isLeaf(root *TreeNode) bool {
	if root.left == nil && root.right == nil {
		return true
	}
	return false
}
