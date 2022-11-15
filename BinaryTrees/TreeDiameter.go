package main

import (
	"fmt"
	"math"
)

func (t *BinaryTree) TreeDiameter(root *TreeNode) {
	var maxDia = 0
	dia := t.Diameter(root, &maxDia)
	fmt.Println(dia)
}
func (t *BinaryTree) Diameter(root *TreeNode, maxDia *int) int {
	if root == nil {
		return 0
	}
	lh := t.Diameter(root.left, maxDia)
	rh := t.Diameter(root.right, maxDia)
	*maxDia = int(math.Max(float64(*maxDia), float64(lh+rh)))
	return 1 + int(math.Max(float64(lh), float64(rh)))
}
