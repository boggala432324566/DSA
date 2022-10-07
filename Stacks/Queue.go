package main

type Queue struct {
	queue []string
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
