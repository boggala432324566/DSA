package main

import (
	"fmt"
	"math"
)

func (g *Graph) BellmanFord(edges []*Node, src *Vertex) {
	length := len(g.Vertexes)
	distance := make([]int, length)
	for i := 0; i < length; i++ {
		distance[i] = math.MaxInt
	}
	distance[src.Key] = 0
	for i := 0; i < len(edges)-1; i++ {
		for _, adj := range edges {
			if distance[adj.from]+adj.Weight < distance[adj.to] {
				distance[adj.to] = distance[adj.from] + adj.Weight
			}
		}
	}
	flag := false
	for _, adj := range edges {
		if distance[adj.from]+adj.Weight < distance[adj.to] {
			distance[adj.to] = distance[adj.from] + adj.Weight
			fmt.Println("Negative Cycle")
			flag = true
		}
	}
	if !flag {
		for i := 0; i < length; i++ {
			fmt.Printf("%d ", distance[i])
		}
	}
}
