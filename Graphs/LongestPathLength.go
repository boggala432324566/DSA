package main

import (
	"fmt"
)

var maxLength int = -1
var length int = 0
var farestNode *Vertex

func (g *Graph) LongestPathLength(v *Vertex) {
	length := len(g.Vertexes)
	visited := make([]bool, length)
	g.LongPath(v, visited)
	fmt.Printf("Long path length:%d ", maxLength)
}
func (g *Graph) LongPath(v *Vertex, visited []bool) {
	visited[v.Key] = true
	for _, adj := range v.Adjacent {
		if !visited[adj.Key] {
			length++
			g.LongPath(adj, visited)
			length--
		}
	}
	if length > maxLength {
		maxLength = length
		farestNode = v
	}
}
