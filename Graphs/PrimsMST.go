package main

import (
	"fmt"
	"math"
)

func (g *Graph) PrimsMST(s *Vertex, n int) {
	mst := make([]bool, n)
	keyVal := make([]int, n)
	parent := make([]int, n)
	for i := 0; i < n; i++ {
		parent[i] = -1
		mst[i] = false
		keyVal[i] = math.MaxInt
	}
	keyVal[s.Key] = 0
	parent[s.Key] = -1
	for i := 0; i < n-1; i++ {
		minVal := math.MaxInt
		var u int
		for v := 0; v < n; v++ {
			if mst[v] == false && minVal > keyVal[v] {
				minVal = keyVal[v]
				u = v
			}
		}
		mst[u] = true
		vertex := g.GetVertex(u)
		for _, adj := range vertex.Adjacent {
			if mst[adj.Key] == false && adj.Weight < keyVal[adj.Key] {
				parent[adj.Key] = u
				keyVal[adj.Key] = adj.Weight
			}
		}
	}
	for i := 1; i < n; i++ {
		fmt.Printf("%d-->%d\n", i, parent[i])
	}
}
