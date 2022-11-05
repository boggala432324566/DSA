package main

import (
	"fmt"
	"math"
)

func (g *Graph) BridgesInGraph(v *Vertex) {
	length := len(g.Vertexes)
	visited := make([]bool, length)
	low := make([]int, length)
	insert := make([]int, length)
	g.PrintBridges(v, g.CreateVertex(-1), low, insert, visited, 0)

}

func (g *Graph) PrintBridges(v *Vertex, parent *Vertex, low []int, insert []int, visited []bool, timer int) {
	timer = timer + 1
	low[v.Key] = timer
	insert[v.Key] = timer
	visited[v.Key] = true
	for _, adj := range v.Adjacent {
		if v.Key == parent.Key {
			continue
		}
		if !visited[adj.Key] {
			g.PrintBridges(adj, v, low, insert, visited, timer)
			low[v.Key] = int(math.Min(float64(low[v.Key]), float64(low[adj.Key])))
			if low[adj.Key] > insert[v.Key] {
				fmt.Printf("%d ---> %d", v.Key, adj.Key)
			}
		} else {
			low[v.Key] = int(math.Min(float64(low[v.Key]), float64(insert[adj.Key])))
		}
	}
}
