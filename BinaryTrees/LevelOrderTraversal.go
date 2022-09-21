package main

import "fmt"

func (t *BinaryTree) LevelOrderTraversal() {
	r := t.root
	t.queue.EnQueue(r)
	fmt.Printf("Level Order Traversal: ")
	for !t.queue.isEmpty() {
		ele := t.queue.DeQueue()
		fmt.Printf("%d ", ele.val)
		if ele.left != nil {
			t.queue.EnQueue(ele.left)
		}
		if ele.right != nil {
			t.queue.EnQueue(ele.right)
		}
	}
}
