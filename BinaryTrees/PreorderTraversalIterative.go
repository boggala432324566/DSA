package main

import "fmt"

func (t *BinaryTree) PreorderTraversalIterative(root *TreeNode) {
	if root == nil {
		return
	}
	t.stack.Push(root)
	for !t.stack.isEmpty() {
		ele := t.stack.Peek()
		fmt.Printf("%d ", ele.val)
		t.stack.Pop()
		if ele.right != nil {
			t.stack.Push(ele.right)
		}
		if ele.left != nil {
			t.stack.Push(ele.left)
		}
	}
}
