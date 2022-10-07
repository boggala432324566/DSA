package main

var s1 Stack
var s2 Stack

func (q *Queue) EnQ(ele string) {
	for !s1.isEmpty() {
		s2.push(s1.peek())
		s1.pop()
	}
	s1.push(ele)
	for !s2.isEmpty() {
		s1.push(s2.peek())
		s2.pop()
	}

}
func (q *Queue) DeQ() string {
	ele := s1.peek()
	s1.pop()
	return ele
}
