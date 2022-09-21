package main

func (l *LinkedList) RotateLinkedListRightByKPlaces(pos int) {
	temp := l.head
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
		temp = first
		first = first.next
		second = second.next
	}
	second.next = l.head
	l.head = temp.next
	temp.next = nil
}
