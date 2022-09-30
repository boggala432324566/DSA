package main

import "fmt"

func (t *BinaryTree) PostorderTraversalIterative(root *TreeNode) {
	if root == nil {
		return
	}
	t.stack.Push(root)
	s2 := Stack{}
	for !t.stack.isEmpty() {
		ele := t.stack.Pop()
		s2.Push(ele)
		if ele.left != nil {
			t.stack.Push(ele.left)
		}
		if ele.right != nil {
			t.stack.Push(ele.right)
		}
	}
	for !s2.isEmpty() {
		ele := s2.Pop()
		fmt.Printf("%d ", ele.val)
	}
}
