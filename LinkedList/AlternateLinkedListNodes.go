package main

func (l *LinkedList) AlternateLinkedListNodes() {
	prev := &Node{}
	slow := l.head
	fast := l.head
	for fast != nil && fast.next != nil {
		prev = slow
		fast = fast.next.next
		slow = slow.next
	}
	if fast != nil {
		prev = slow
		slow = slow.next
	}
	prev.next = nil
	secondHead := l.reverseLinkedList(slow)
	l.alternateNodes(l.head, secondHead)
}

func (l *LinkedList) alternateNodes(head1 *Node, head2 *Node) {
	for head1 != nil && head2 != nil {
		temp1 := head1.next
		temp2 := head2.next
		head1.next = head2
		head2.next = temp1
		head1 = temp1
		head2 = temp2
	}
}
