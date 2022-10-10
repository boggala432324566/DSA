package main

import "fmt"

var cyclePresent bool = false

func (g *Graph) CycleInDirectedGraph(v *Vertex) {
	length := len(g.Vertexes)
	visited := make([]bool, length)
	inPath := make([]bool, length)
	g.Cycle(v, visited, inPath)
	fmt.Printf("Cycle in graph:%v ", cyclePresent)
}

func (g *Graph) Cycle(v *Vertex, visited []bool, inPath []bool) {
	visited[v.Key] = true
	inPath[v.Key] = true
	for _, adj := range v.Adjacent {
		if inPath[adj.Key] {
			cyclePresent = true
			return
		} else if !visited[adj.Key] {
			g.Cycle(adj, visited, inPath)
		}
	}
	inPath[v.Key] = false
}
