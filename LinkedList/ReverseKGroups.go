package main

func (l *LinkedList) reverseList() {
	prev := &Node{}
	prev = nil
	curr := l.head
	nxt := l.head.next

	for nxt != nil {
		curr.next = prev
		prev = curr
		curr = nxt
		nxt = nxt.next
	}
	curr.next = prev
	l.head = curr
}
