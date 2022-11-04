package main

import "fmt"

type Vertex struct {
	Key      int
	Adjacent []*Vertex
	Weight   int
}
type Graph struct {
	Vertexes   []*Vertex
	GraphQueue GraphQueue
	GraphStack GraphStack
}

type Node struct {
	from   int
	to     int
	Weight int
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
func (g *Graph) AddEdge(fromVertex int, toVertex int, edgeWeight int, tr bool) *Node {
	fv := g.GetVertex(fromVertex)
	tv := g.GetVertex(toVertex)
	if fv == nil || tv == nil {
		fmt.Errorf("invalid Vertex")
		return nil
	} else if g.Contains(fv.Adjacent, tv.Key) {
		fmt.Errorf("duplicate edge")
		return nil
	} else if !tr {
		tv.Weight = edgeWeight
		fv.Adjacent = append(fv.Adjacent, tv)
	} else {
		tv.Weight = edgeWeight
		tv.Adjacent = append(tv.Adjacent, fv)
	}
	return &Node{from: fv.Key, to: tv.Key, Weight: edgeWeight}
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
	g.AddVertex(4)
	var edges []*Node
	edges = append(edges, g.AddEdge(0, 1, 2, false))
	edges = append(edges, g.AddEdge(0, 4, 4, false))
	edges = append(edges, g.AddEdge(1, 4, 1, false))
	edges = append(edges, g.AddEdge(1, 2, 3, false))
	edges = append(edges, g.AddEdge(2, 3, 4, false))
	edges = append(edges, g.AddEdge(3, 4, 5, false))

	/*g.AddEdge(0, 1, 1, false)
	g.AddEdge(4, 0, 1, false)
	g.AddEdge(1, 4, 1, false)
	g.AddEdge(1, 2, 1, false)
	g.AddEdge(2, 3, 1, false)
	g.AddEdge(3, 4, 1, false)*/
	//g.Print()
	//g.BFS(g.Vertexes[2])
	//g.DFSUsingIterativeApproach(g.Vertexes[0])
	//g.DFSUsingRecursion(g.Vertexes[0])
	//g.IsLands()
	//g.LongestPathLength(g.Vertexes[0])
	//g.CycleInDirectedGraph(g.Vertexes[0])
	//g.PrimsMST(g.Vertexes[0], 5)
	//g.KruskalsMST(edges, 5)
	//g.FloydWarshall()
	//g.ShortestPath(g.Vertexes[0])
	//g.BipartiteGraph(g.Vertexes[0])
	//g.KosaRaju(g.Vertexes[0])
	//g.DijkstraShortestPath(g.Vertexes[0])
	g.BellmanFord(edges, g.Vertexes[0])
}
