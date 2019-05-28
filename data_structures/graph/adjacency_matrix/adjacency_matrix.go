package main

import (
	"fmt"
)

// graph represents an adjacency matrix graph.
type graph struct {
	matrix [][]int
	size   int
}

// newGraph returns a new graph with the given size.
func newGraph(size int) *graph {
	matrix := make([][]int, size)
	for i := range matrix {
		matrix[i] = make([]int, size)
	}
	return &graph{matrix: matrix, size: size}
}

// addEdge adds an edge to the graph from -> to.
func (g *graph) addEdge(from, to int) {
	if from < 0 || to < 0 || from >= g.size || to >= g.size {
		return
	}
	g.matrix[from][to] = 1
}

// print prints the adjacency matrix.
func (g *graph) print() {
	for _, row := range g.matrix {
		fmt.Println(row)
	}
}

func main() {
	size := 5
	graph := newGraph(size)

	graph.addEdge(0, 1)
	graph.addEdge(0, 3)
	graph.addEdge(1, 2)
	graph.addEdge(2, 4)
	graph.addEdge(3, 4)

	graph.print()
}
