package main

import "fmt"

func (g *Graph) DFSUsingRecursion(v *Vertex) {
	length := len(g.Vertexes)
	visited := make([]bool, length)
	g.DFS(v, visited)

}

func (g *Graph) DFS(v *Vertex, visited []bool) {
	visited[v.Key] = true
	fmt.Printf("%d ", v.Key)
	for _, adj := range v.Adjacent {
		if !visited[adj.Key] {
			g.DFS(adj, visited)
		}
	}
}
