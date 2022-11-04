package main

import (
	"fmt"
	"math"
)

func (g *Graph) DijkstraShortestPath(v *Vertex) {
	length := len(g.Vertexes)
	distance := make([]int, length)
	for i := 0; i < length; i++ {
		distance[i] = math.MaxInt
	}
	g.GraphQueue.EnQueue(v)
	distance[v.Key] = 0
	for !g.GraphQueue.isEmpty() {
		ele := g.GraphQueue.DeQueueMin()
		for _, adj := range ele.Adjacent {
			if distance[ele.Key]+adj.Weight < distance[adj.Key] {
				distance[adj.Key] = distance[ele.Key] + adj.Weight
				g.GraphQueue.EnQueue(adj)
			}
		}
	}
	for i := 0; i < length; i++ {
		fmt.Printf("%d ", distance[i])
	}
}
