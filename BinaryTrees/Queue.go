package main

type Queue struct {
	queue []*TreeNode
}

func (q *Queue) EnQueue(node *TreeNode) {
	q.queue = append(q.queue, node)
}

func (q *Queue) DeQueue() *TreeNode {
	tNode := q.queue[0]
	q.queue = q.queue[1:]
	return tNode
}

func (q *Queue) Peek() *TreeNode {
	return q.queue[0]

}

func (q *Queue) isEmpty() bool {
	return len(q.queue) == 0
}
