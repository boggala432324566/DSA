package main

import (
	"fmt"
)

func (g *Graph) BipartiteGraph(v *Vertex) {
	length := len(g.Vertexes)
	color := make([]int, length)
	for i := 0; i < length; i++ {
		color[i] = -1
	}
	g.GraphQueue.EnQueue(v)
	color[v.Key] = 1
	for !g.GraphQueue.isEmpty() {
		ele := g.GraphQueue.Peek()
		g.GraphQueue.DeQueue()
		for _, adj := range ele.Adjacent {
			if color[adj.Key] == -1 {
				color[adj.Key] = 1 - color[ele.Key]
				g.GraphQueue.EnQueue(adj)
			} else if color[ele.Key] == color[adj.Key] {
				fmt.Printf("Not Bipartite")
				return
			}
		}
	}
	fmt.Printf("Bipartite")
}
