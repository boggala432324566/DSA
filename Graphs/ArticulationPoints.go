package main

import (
	"fmt"
	"math"
)

func (g *Graph) ArticulationPoints(v *Vertex) {
	length := len(g.Vertexes)
	visited := make([]bool, length)
	low := make([]int, length)
	insert := make([]int, length)
	articulation := make([]int, length)
	g.Articulation(v, g.CreateVertex(-1), low, insert, visited, articulation, 0)
	for i := 0; i < length; i++ {
		if articulation[i] == 1 {
			fmt.Println(articulation[i])
		}
	}

}

func (g *Graph) Articulation(v *Vertex, parent *Vertex, low []int, insert []int, visited []bool, articulation []int, timer int) {
	timer = timer + 1
	low[v.Key] = timer
	insert[v.Key] = timer
	visited[v.Key] = true
	child := 0
	for _, adj := range v.Adjacent {
		if v.Key == parent.Key {
			continue
		}
		if !visited[adj.Key] {
			g.Articulation(adj, v, low, insert, visited, articulation, timer)
			low[v.Key] = int(math.Min(float64(low[v.Key]), float64(low[adj.Key])))
			if low[adj.Key] >= insert[v.Key] && parent.Key != -1 {
				articulation[v.Key] = 1
			}
			child++
		} else {
			low[v.Key] = int(math.Min(float64(low[v.Key]), float64(insert[adj.Key])))
		}
	}
	if child > 1 && parent.Key != -1 {
		articulation[v.Key] = 1
	}
}
