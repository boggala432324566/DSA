package main

import "fmt"

func (g *Graph) CycleInUnDirectedGraph(v *Vertex) {
	length := len(g.Vertexes)
	visited := make([]bool, length)
	parent := make([]int, length)
	for i := 0; i < length; i++ {
		parent[i] = -1
	}
	cyclePresent := g.IsCyclic(v, visited, parent)
	fmt.Printf("Cycle in graph:%v ", cyclePresent)
}
func (g *Graph) IsCyclic(v *Vertex, visited []bool, parent []int) bool {
	g.GraphQueue.EnQueue(v)
	visited[v.Key] = true
	for !g.GraphQueue.isEmpty() {
		ele := g.GraphQueue.Peek()
		g.GraphQueue.DeQueue()
		for _, adj := range ele.Adjacent {
			if !visited[adj.Key] {
				visited[adj.Key] = true
				parent[adj.Key] = ele.Key
				g.GraphQueue.EnQueue(adj)
			} else if parent[ele.Key] != adj.Key {
				return true
			}
		}
	}
	return false
}
