package main

func (l *LinkedList) AlternateLinkedListNodes() {
	first := l.head
	second := l.head
	for second.next == nil {
		temp := &Node{}
		for second.next == nil {
			temp = second
			second = second.next
		}
		temp.next = nil
		second.next = first.next
		first.next = second
		first = second.next
		second = second.next
	}
}
