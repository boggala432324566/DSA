package main

type GraphQueue struct {
	queue []*Vertex
}

func (q *GraphQueue) EnQueue(v *Vertex) {
	q.queue = append(q.queue, v)
}
func (q *GraphQueue) DeQueue() *Vertex {
	ele := q.queue[0]
	q.queue = q.queue[1:]
	return ele
}

func (q *GraphQueue) Peek() *Vertex {
	return q.queue[0]

}

func (q *GraphQueue) isEmpty() bool {
	return len(q.queue) == 0
}

func (q *GraphQueue) Size() int {
	return len(q.queue)
}
