// See: https://www.youtube.com/watch?v=JDP1OVgoa0Q
// See: https://www.youtube.com/watch?v=bSZ57h7GN2w
package main

import "fmt"

// Graph represents an adjacency list graph.
type graph struct {
	vertices []*vertex
}

// Vertex represents a graph vertex.
type vertex struct {
	key int
	adj []*vertex
}

// addVertex adds a vertext to the graph.
func (g *graph) addVertex(k int) {
	if contains(g.vertices, k) {
		fmt.Println(fmt.Errorf("Vertex %v not added because it is an existing key", k))
		return
	}
	g.vertices = append(g.vertices, &vertex{key: k})
}

// addEdge adds an edge to the graph.
func (g *graph) addEdge(from, to int) {
	// get vertex
	fromVertex := g.getVertex(from)
	toVertex := g.getVertex(to)

	// check error
	if fromVertex == nil || toVertex == nil {
		fmt.Println("Invalid edge")
		return
	}

	// check if edge already exists
	if contains(fromVertex.adj, to) {
		fmt.Println("Existing edge")
		return
	}

	// add edge
	fromVertex.adj = append(fromVertex.adj, toVertex)
}

// getVertex returns a pointer to the vertex.
func (g *graph) getVertex(k int) *vertex {
	for _, v := range g.vertices {
		if k == v.key {
			return v
		}
	}
	return nil
}

// contains returns true if the key exists in the slice.
func contains(s []*vertex, k int) bool {
	for _, v := range s {
		if k == v.key {
			return true
		}
	}
	return false
}

// print prints the adjacency list.
func (g *graph) print() {
	for _, v := range g.vertices {
		fmt.Printf("\nVertex %v : ", v.key)
		for _, v := range v.adj {
			fmt.Printf("%v ", v.key)
		}
	}
}

func main() {
	g := &graph{}
	for i := 0; i < 5; i++ {
		g.addVertex(i)
	}
	g.addEdge(1, 2)
	g.addEdge(1, 3)
	g.addEdge(2, 3)
	g.addEdge(4, 1)
	g.addEdge(4, 2)
	g.addEdge(4, 3)
	g.print()
}
