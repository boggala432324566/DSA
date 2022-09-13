package main

func (l *LinkedList) DeleteGivenNodeFromEnd(pos int) {
	first := l.head
	second := l.head
	for i := 1; i < pos; i++ {
		second = second.next
	}
	if second.next == nil {
		l.head = first.next
		return
	}
	for second.next != nil {
		first = first.next
		second = second.next
	}
	first.val = first.next.val
	first.next = first.next.next

}
