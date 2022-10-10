package main

func (l *LinkedList) reverseByGroupOfK(head *Node, length int, k int) {
	if length < k || head == nil {
		return
	}
	dummy := &Node{}
	dummy.next = head
	nxt := dummy
	prev := dummy
	curr := dummy
	for length >= k {
		curr = prev.next
		nxt = curr.next
		for i := 1; i < k; i++ {
			curr.next = nxt.next
			nxt.next = prev.next
			prev.next = nxt
			nxt = curr.next
		}
		prev = curr
		length = length - k
	}
	l.head = dummy.next
}
