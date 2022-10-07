package main

import "fmt"

func (g *Graph) BFS(v *Vertex) {
	length := len(g.Vertexes)
	visited := make([]bool, length)
	g.GraphQueue.EnQueue(v)
	visited[v.Key] = true
	for !g.GraphQueue.isEmpty() {
		ele := g.GraphQueue.Peek()
		g.GraphQueue.DeQueue()
		fmt.Printf("%d ", ele.Key)
		for _, adj := range ele.Adjacent {
			if !visited[adj.Key] {
				visited[adj.Key] = true
				g.GraphQueue.EnQueue(adj)
			}
		}
	}
}
