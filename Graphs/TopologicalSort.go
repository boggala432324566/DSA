package main

import "fmt"

func (g *Graph) TopologicalSort(v *Vertex) {
	length := len(g.Vertexes)
	visited := make([]bool, length)
	g.TopoSort(v, visited, g.GraphStack)
	fmt.Printf("Topological Sort:")
	if g.GraphStack.isEmpty() {
		fmt.Printf("%d ", g.GraphStack.pop())
	}
}

func (g *Graph) TopoSort(v *Vertex, visited []bool, st GraphStack) {
	visited[v.Key] = true
	for _, adj := range v.Adjacent {
		if !visited[adj.Key] {
			g.TopoSort(adj, visited, st)
		}
	}
	st.push(v)
}
