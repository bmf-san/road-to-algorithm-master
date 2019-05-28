package main

import "testing"

func TestNewGratph(t *testing.T) {
	size := 5
	g := newGraph(size)
	if g.size != size {
		t.Error("Size not set")
	}
}

func TestAddEdge(t *testing.T) {
	size := 5
	g := newGraph(size)
	g.addEdge(0, 1)
	if g.matrix[0][1] != 1 {
		t.Error("Edge not added")
	}
}
