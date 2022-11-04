package main

import "fmt"

func (g *Graph) FloydWarshall() {
	graph := [][]int{{0, 4, 0, 0, 2}, {4, 0, 5, 0, 7}, {0, 5, 0, 6, 0}, {0, 0, 6, 0, 3}, {2, 7, 0, 3, 0}}
	g.FloydWar(graph, 5)
}

func (g *Graph) FloydWar(graph [][]int, n int) {
	dist := make([][]int, n)
	for i := 0; i < n; i++ {
		dist[i] = graph[i]
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			for k := 0; k < n; k++ {
				if dist[i][k]+dist[k][j] < dist[i][j] {
					dist[i][j] = dist[i][k] + dist[k][j]
				}
			}
		}
	}
	for i := 0; i < n; i++ {
		fmt.Println(dist[i])
	}
}
