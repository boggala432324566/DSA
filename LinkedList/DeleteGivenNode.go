package main

func (l *LinkedList) DeleteGivenNode(node *Node) {
	node.val = node.next.val
	node.next = node.next.next
}
