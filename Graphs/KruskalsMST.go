package main

import (
	"fmt"
	"sort"
)

func FindPar(v int, parent []int) int {
	if v == parent[v] {
		return v
	}
	return FindPar(parent[v], parent)
}

func Union(u int, v int, parent []int, rank []int) {
	u = FindPar(u, parent)
	v = FindPar(v, parent)
	if rank[u] < rank[v] {
		parent[u] = v
	} else if rank[v] < rank[u] {
		parent[v] = u
	} else {
		parent[v] = u
		rank[u]++
	}

}
func (g *Graph) KruskalsMST(adj []*Node, n int) {
	sort.Slice(adj, func(i, j int) bool {
		return adj[i].Weight < adj[j].Weight
	})
	parent := make([]int, n)
	rank := make([]int, n)
	for i := 0; i < n; i++ {
		parent[i] = i
		rank[i] = 0
	}
	var mst []*Node
	mstCost := 0
	for _, v := range adj {
		if FindPar(v.from, parent) != FindPar(v.to, parent) {
			mstCost += v.Weight
			mst = append(mst, v)
			Union(v.from, v.to, parent, rank)
		}
	}
	fmt.Printf("MST Cost:%d\n", mstCost)
	for _, v := range mst {
		fmt.Printf("%d-->%d\n", v.from, v.to)
	}
}
