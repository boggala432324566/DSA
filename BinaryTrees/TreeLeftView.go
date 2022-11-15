package main

import (
	"fmt"
)

func (t *BinaryTree) TreeLeftView(root *TreeNode) {
	var slice []int
	t.LeftView(&slice, root, 0)
	fmt.Println("Left View:")
	for _, v := range slice {
		fmt.Println(v)
	}

}

func (t *BinaryTree) LeftView(slice *[]int, root *TreeNode, level int) {
	if root == nil {
		return
	}
	if level == len(*slice) {
		*slice = append(*slice, root.val)
	}
	t.LeftView(slice, root.left, level+1)
	t.LeftView(slice, root.right, level+1)
}
