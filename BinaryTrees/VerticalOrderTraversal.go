package main

import (
	"fmt"
	"sort"
)

func (t *BinaryTree) VerticalOrderTraversal(root *TreeNode) {
	var vOrderTraversal = make(map[int][]int)
	t.VerticalOrder(vOrderTraversal, root, 0)
	var distSlice []int
	for d := range vOrderTraversal {
		distSlice = append(distSlice, d)
	}
	sort.Ints(distSlice)
	for _, d := range distSlice {
		fmt.Printf("distance[%d]:", d)
		fmt.Println(vOrderTraversal[d])
	}

}

func (t *BinaryTree) VerticalOrder(vOrderTraversal map[int][]int, root *TreeNode, distance int) {
	if root == nil {
		return
	}
	vOrderTraversal[distance] = append(vOrderTraversal[distance], root.val)
	t.VerticalOrder(vOrderTraversal, root.left, distance-1)
	t.VerticalOrder(vOrderTraversal, root.right, distance+1)
}
