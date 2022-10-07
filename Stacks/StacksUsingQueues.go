package main

var q1 Queue
var q2 Queue

func (s *Stack) StackPush(ele string) {
	q1.EnQueue(ele)
	for !q2.isEmpty() {
		q1.EnQueue(q2.Peek())
		q1.DeQueue()
	}
	for !q1.isEmpty() {
		q2.EnQueue(q1.Peek())
		q1.DeQueue()
	}
}
func (s *Stack) StackPop() string {
	ele := q2.DeQueue()
	return ele
}
