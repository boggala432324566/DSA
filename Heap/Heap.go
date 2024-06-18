// You can edit this code!
// Click here and start typing.
package main

import (
	"container/heap"
	"fmt"
)

type queue struct {
	V int
	P int
}

type pq []queue

func (p pq) Len() int {
	return len(p)
}

func (p pq) Less(i, j int) bool {
	return p[i].P > p[j].P
}

func (p pq) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]

}

func (p *pq) Push(x interface{}) {
	item := x.(queue)
	*p = append(*p, item)
}

func (p *pq) Pop() interface{} {
	old := *p
	n := len(old)
	item := old[n-1]
	*p = old[0 : n-1]
	return item

}

func main() {
	pqueue := &pq{}
	pqueue.Push(queue{V: 5, P: 2})
	pqueue.Push(queue{V: 6, P: 1})
	pqueue.Push(queue{V: 2, P: 9})
	pqueue.Push(queue{V: 7, P: 4})
	pqueue.Push(queue{V: 9, P: 5})
	pqueue.Push(queue{V: 8, P: 3})
	heap.Init(pqueue)
	heap.Push(pqueue, queue{V: 1, P: 10})
	for pqueue.Len() != 0 {
		fmt.Println(heap.Pop(pqueue))
	}

}
