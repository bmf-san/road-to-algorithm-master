package main

import "testing"

func TestAddVertex(t *testing.T) {
	g := &graph{}
	g.addVertex(1)
	if g.vertices[0].key != 1 {
		t.Error("Vertex not added")
	}
}

func TestAddEdge(t *testing.T) {
	g := &graph{}
	g.addVertex(1)
	g.addVertex(2)
	g.addEdge(1, 2)
	if g.vertices[0].adj[0].key != 2 {
		t.Error("Edge not added")
	}
}

func TestGetVertex(t *testing.T) {
	g := &graph{}
	g.addVertex(1)
	v := g.getVertex(1)
	if v.key != 1 {
		t.Error("Vertex not returned")
	}
}

func TestContains(t *testing.T) {
	g := &graph{}
	g.addVertex(1)
	g.addVertex(2)
	if !contains(g.vertices, 2) {
		t.Error("Vertex not found")
	}
}
