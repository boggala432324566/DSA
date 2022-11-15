package main

import (
	"fmt"
	"sort"
)

func (t *BinaryTree) TreeButtomView(root *TreeNode) {
	var vOrderTraversal = make(map[int]int)
	t.ButtomView(vOrderTraversal, root, 0)
	var distSlice []int
	for d := range vOrderTraversal {
		distSlice = append(distSlice, d)
	}
	sort.Ints(distSlice)
	fmt.Println("Buttom View:")
	for _, d := range distSlice {
		fmt.Println(vOrderTraversal[d])
	}

}

func (t *BinaryTree) ButtomView(vOrderTraversal map[int]int, root *TreeNode, distance int) {
	if root == nil {
		return
	}
	vOrderTraversal[distance] = root.val
	t.ButtomView(vOrderTraversal, root.left, distance-1)
	t.ButtomView(vOrderTraversal, root.right, distance+1)
}
