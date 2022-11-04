package main

import "math"

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

func (q *GraphQueue) DeQueueMin() *Vertex {
	min := math.MaxInt
	minInd := 0
	for i, val := range q.queue {
		if val.Key < min {
			min = val.Weight
			minInd = i
		}
	}
	ele := q.queue[minInd]
	q.queue = append(q.queue[:minInd], q.queue[minInd+1:]...)
	return ele
}
