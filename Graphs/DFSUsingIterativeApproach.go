package main

import "fmt"

func (g *Graph) DFSUsingIterativeApproach(v *Vertex) {
	length := len(g.Vertexes)
	visited := make([]bool, length)
	g.GraphStack.push(v)
	visited[v.Key] = true
	for !g.GraphStack.isEmpty() {
		ele := g.GraphStack.peek()
		g.GraphStack.pop()
		fmt.Printf("%d ", ele.Key)
		for _, adj := range ele.Adjacent {
			if !visited[adj.Key] {
				visited[adj.Key] = true
				g.GraphStack.push(adj)
			}
		}
	}
}
