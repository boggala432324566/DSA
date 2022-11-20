package main

import (
	"fmt"
	"math"
)

func (t *BST) BSTOrNot() {
	bstFlag := CheckBST(t.root, math.MinInt, math.MaxInt)
	if bstFlag {
		fmt.Println("BST")
	} else {
		fmt.Println("Not BST")
	}
}

func CheckBST(root *BSTNode, minInt int, maxInt int) bool {
	if root == nil {
		return true
	}
	if root.val < minInt || root.val > maxInt {
		return false
	}
	return CheckBST(root.left, root.val, maxInt) && CheckBST(root.right, minInt, root.val)

}
