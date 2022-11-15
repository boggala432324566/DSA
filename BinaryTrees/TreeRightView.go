package main

import (
	"fmt"
)

func (t *BinaryTree) TreeRightView(root *TreeNode) {
	var slice []int
	t.RightView(&slice, root, 0)
	fmt.Println("Right View:")
	for _, v := range slice {
		fmt.Println(v)
	}

}

func (t *BinaryTree) RightView(slice *[]int, root *TreeNode, level int) {
	if root == nil {
		return
	}
	if level == len(*slice) {
		*slice = append(*slice, root.val)
	}
	t.RightView(slice, root.right, level+1)
	t.RightView(slice, root.left, level+1)
}
