package main

import (
	"fmt"
	"math"
)

func (g *Graph) ShortestPath(v *Vertex) {
	length := len(g.Vertexes)
	distance := make([]int, length)
	for i := 0; i < length; i++ {
		distance[i] = math.MaxInt
	}
	g.GraphQueue.EnQueue(v)
	distance[v.Key] = 0
	for !g.GraphQueue.isEmpty() {
		ele := g.GraphQueue.Peek()
		g.GraphQueue.DeQueue()
		for _, adj := range ele.Adjacent {
			if distance[ele.Key]+1 < distance[adj.Key] {
				distance[adj.Key] = distance[ele.Key] + 1
				g.GraphQueue.EnQueue(adj)
			}
		}
	}
	for i := 0; i < length; i++ {
		fmt.Printf("%d ", distance[i])
	}
}
