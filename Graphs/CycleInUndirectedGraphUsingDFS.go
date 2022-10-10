package main

import "fmt"

var cyclePre bool = false

func (g *Graph) CycleInUndirectedGraphUsingDFS(v *Vertex) {
	length := len(g.Vertexes)
	visited := make([]bool, length)
	parent := make([]int, length)
	for i := 0; i < length; i++ {
		parent[i] = -1
	}
	g.CycleUsingDFS(v, visited, parent)
	fmt.Printf("Cycle in graph:%v ", cyclePresent)
}

func (g *Graph) CycleUsingDFS(v *Vertex, visited []bool, parent []int) {
	visited[v.Key] = true
	for _, adj := range v.Adjacent {
		if !visited[adj.Key] {
			parent[adj.Key] = v.Key
			g.CycleUsingDFS(adj, visited, parent)
		} else if parent[v.Key] != adj.Key {
			cyclePre = true
			return
		}
	}

}
