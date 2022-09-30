package main

import "fmt"

func (t *BinaryTree) InorderTraversalIterative(root *TreeNode) {
	if root == nil {
		return
	}
	curr := root
	for curr != nil || !t.stack.isEmpty() {
		for curr != nil {
			t.stack.Push(curr)
			curr = curr.left
		}
		ele := t.stack.Peek()
		fmt.Printf("%d ", ele.val)
		t.stack.Pop()
		curr = ele.right
	}
}
