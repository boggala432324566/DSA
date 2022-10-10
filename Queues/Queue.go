package main

import (
	"fmt"
)

type Queue struct {
	queue []string
	Stack Stack
}

func (q *Queue) EnQueue(ele string) {
	q.queue = append(q.queue, ele)
}
func (q *Queue) DeQueue() string {
	ele := q.queue[0]
	q.queue = q.queue[1:]
	return ele
}

func (q *Queue) Peek() string {
	return q.queue[0]

}

func (q *Queue) isEmpty() bool {
	return len(q.queue) == 0
}

func (q *Queue) Size() int {
	return len(q.queue)
}

func main() {
	q := Queue{}
	q.EnQ("23")
	q.EnQ("24")
	q.EnQ("45")
	fmt.Println(q.DeQ())
	//q.SmallestNumberDividesDigits0ans9(111)
}
