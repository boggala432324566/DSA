package main

func (l *LinkedList) reverseByGroupOfK(head *Node, length int, k int) *Node {
	if length < k {
		return head
	}
	prev := &Node{}
	nxt := &Node{}
	prev = nil
	nxt = nil
	curr := l.head
	count := 0
	for nxt != nil && count < k {
		nxt = curr.next
		curr.next = prev
		prev = curr
		curr = nxt
		count++
	}
	if nxt != nil {
		head.next = l.reverseByGroupOfK(nxt, length-k, k)
	}
	return prev
}
