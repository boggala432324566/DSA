package main

import "fmt"

func TopoSort(v *Vertex, visited []bool, st *GraphStack) {
	visited[v.Key] = true
	for _, adj := range v.Adjacent {
		if !visited[adj.Key] {
			TopoSort(adj, visited, st)
		}
	}
	st.push(v)
}

func DFS(v *Vertex, visited []bool) {
	visited[v.Key] = true
	fmt.Printf("%d ", v.Key)
	for _, adj := range v.Adjacent {
		if !visited[adj.Key] {
			DFS(adj, visited)
		}
	}
}
func (g *Graph) KosaRaju(v *Vertex) {
	st := GraphStack{}
	length := len(g.Vertexes)
	visited := make([]bool, length)
	TopoSort(v, visited, &st)
	for i := 0; i < length; i++ {
		visited[i] = false
		v := g.GetVertex(i)
		for _, adj := range v.Adjacent {
			g.AddEdge(v.Key, adj.Key, 1, true)
		}
	}
	for !st.isEmpty() {
		ele := st.pop()
		if visited[ele.Key] == false {
			fmt.Println("SCC:")
			DFS(ele, visited)
			fmt.Println()
		}
	}
}
