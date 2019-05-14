package main

import (
	"errors"
	"fmt"
)

// Edge is the structure representing an edge.
type Edge struct {
	To     int
	Weight int
}

// Vertex is the structure representing a vertex.
type Vertex struct {
	ID        int
	Neighbors []*Edge
}

// Graph is the structure representing a graph.
type Graph struct {
	directed bool
	Vertex   map[int]*Vertex
}

// NewGraph creates a new Graph.
func NewGraph(directed bool) *Graph {
	return &Graph{directed, map[int]*Vertex{}}
}

// AddVertex adds a vertex to the graph.
func (g *Graph) AddVertex(id int) error {
	if g.Vertex[id] != nil {
		return errors.New("vertex already in graph")
	}
	g.Vertex[id] = &Vertex{id, []*Edge{}}
	return nil
}

// AddEdge adds an edge between two vertices with a given weight.
func (g *Graph) AddEdge(from, to, weight int) {
	if g.Vertex[from] == nil {
		g.AddVertex(from)
	}
	if g.Vertex[to] == nil {
		g.AddVertex(to)
	}
	g.Vertex[from].Neighbors = append(g.Vertex[from].Neighbors, &Edge{to, weight})
	// If graph is undirected, add the reverse edge as well.
	if !g.directed {
		g.Vertex[to].Neighbors = append(g.Vertex[to].Neighbors, &Edge{from, weight})
	}
}

// VisitAll counts the number of vertices that can be visited from a vertex v.
func (g *Graph) VisitAll(v int) int {
	s := NewStack()
	visited := map[int]bool{}
	s.Push(v)
	numVisited := 0

	// Depth-first search
	for s.Size() > 0 {
		u := s.Pop()
		if !visited[u] {
			visited[u] = true
			numVisited++
			for _, n := range g.Vertex[u].Neighbors {
				s.Push(n.To)
			}
		}
	}

	return numVisited
}

// IsConnected checks if the graph is connected, i.e. that all vertices can be visted.
func (g *Graph) IsConnected() bool {
	if len(g.Vertex) == 0 {
		return true
	}
	for k := range g.Vertex {
		numVisited := g.VisitAll(k)
		return numVisited == len(g.Vertex)
	}
	// Code should never reach this point.
	return false
}

// ShortestPath returns the length of the shortest path between vertices from and to.
func (g *Graph) ShortestPath(from, to int) (int, error) {
	if g.Vertex[from] == nil {
		return 0, errors.New("from vertex is not in graph")
	}
	if g.Vertex[to] == nil {
		return 0, errors.New("to vertex is not in graph")
	}
	if from == to {
		return 0, nil
	}

	// shortestDist initialized to max value of int.
	intMax := int(^uint(0) >> 1)

	pq := NewDistQueue()
	visited := map[int]bool{}

	dist := map[int]int{}

	for _, v := range g.Vertex {
		var initDist int
		if v.ID == from {
			initDist = 0
		} else {
			initDist = intMax
		}
		dist[v.ID] = initDist
		pq.Enqueue(v.ID, initDist)
	}

	// Dijkstra
	for pq.Size() > 0 {
		u, udist := pq.Dequeue()
		for _, v := range g.Vertex[u].Neighbors {
			if !visited[v.To] {
				visited[v.To] = true
				newDist := udist + v.Weight
				if newDist < dist[v.To] {
					dist[v.To] = newDist
					pq.Update(v.To, newDist)
				}
			}
		}
	}

	return dist[to], nil
}

func main() {
	g := NewGraph(false)
	g.AddEdge(1, 2, 7)
	g.AddEdge(1, 3, 9)
	g.AddEdge(1, 6, 14)
	g.AddEdge(2, 3, 10)
	g.AddEdge(2, 4, 15)
	g.AddEdge(3, 4, 11)
	g.AddEdge(3, 6, 2)
	g.AddEdge(4, 5, 6)
	g.AddEdge(5, 6, 9)

	dist, _ := g.ShortestPath(2, 5)
	fmt.Println(dist)
}
