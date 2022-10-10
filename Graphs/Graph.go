package main

import "fmt"

type Vertex struct {
	Key      int
	Adjacent []*Vertex
}
type Graph struct {
	Vertexes   []*Vertex
	GraphQueue GraphQueue
	GraphStack GraphStack
}

func (g *Graph) CreateVertex(key int) *Vertex {
	return &Vertex{Key: key}
}

func (g *Graph) AddVertex(key int) {
	v := g.CreateVertex(key)
	if !g.Contains(g.Vertexes, v.Key) {
		g.Vertexes = append(g.Vertexes, v)
	}
}
func (g *Graph) AddEdge(fromVertex int, toVertex int) {
	fv := g.GetVertex(fromVertex)
	tv := g.GetVertex(toVertex)
	if fv == nil || tv == nil {
		fmt.Errorf("invalid Vertex")
	} else if g.Contains(fv.Adjacent, tv.Key) {
		fmt.Errorf("duplicate edge")
	} else {
		fv.Adjacent = append(fv.Adjacent, tv)
	}
}

func (g *Graph) GetVertex(key int) *Vertex {
	for _, v := range g.Vertexes {
		if key == v.Key {
			return v
		}
	}
	return nil
}

func (g *Graph) Print() {
	for _, v := range g.Vertexes {
		fmt.Printf("Vertex %d:", v.Key)
		for _, av := range v.Adjacent {
			fmt.Printf(" %d ", av.Key)
		}
		fmt.Println()
	}
}

func (g *Graph) Contains(vList []*Vertex, key int) bool {
	for _, v := range vList {
		if v.Key == key {
			return true
		}
	}
	return false
}

func main() {
	g := Graph{}
	g.AddVertex(0)
	g.AddVertex(1)
	g.AddVertex(2)
	g.AddVertex(3)
	g.AddEdge(0, 2)
	g.AddEdge(0, 1)
	g.AddEdge(1, 3)
	g.AddEdge(3, 2)
	g.AddEdge(2, 1)
	//g.Print()
	//g.BFS(g.Vertexes[2])
	//g.DFSUsingIterativeApproach(g.Vertexes[0])
	//g.DFSUsingRecursion(g.Vertexes[0])
	//g.IsLands()
	//g.LongestPathLength(g.Vertexes[0])
	g.CycleInDirectedGraph(g.Vertexes[0])
}
