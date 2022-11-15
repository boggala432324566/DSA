package main

import (
	"fmt"
	"sort"
)

func (t *BinaryTree) TreeTopView(root *TreeNode) {
	var vOrderTraversal = make(map[int][]int)
	t.TopView(vOrderTraversal, root, 0)
	var distSlice []int
	for d := range vOrderTraversal {
		distSlice = append(distSlice, d)
	}
	sort.Ints(distSlice)
	fmt.Println("Top View:")
	for _, d := range distSlice {
		fmt.Println(vOrderTraversal[d])
	}

}

func (t *BinaryTree) TopView(vOrderTraversal map[int][]int, root *TreeNode, distance int) {
	if root == nil {
		return
	}
	if _, ok := vOrderTraversal[distance]; !ok {
		vOrderTraversal[distance] = append(vOrderTraversal[distance], root.val)
	}
	t.TopView(vOrderTraversal, root.left, distance-1)
	t.TopView(vOrderTraversal, root.right, distance+1)
}
